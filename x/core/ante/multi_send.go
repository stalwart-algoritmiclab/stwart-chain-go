/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package ante

import (
	"math"
	"strconv"

	sdkioerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
)

func (f CoreDecorator) processMsgMultiSend(ctx sdk.Context, tx sdk.Tx, msgSend sdk.Msg, i int) (sdk.Context, error) {
	msg, ok := msgSend.(*types.MsgMultiSend)
	if !ok {
		addr := ""
		if len(msg.Inputs) != 0 {
			addr = msg.Inputs[0].Address
		}
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unexpected MsgMultiSend type %s", addr)
	}

	if len(msg.Inputs) == 0 {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unexpected MsgMultiSend inputs %s", msg.Inputs)
	}

	fromAddress := msg.Inputs[0].Address

	feePayer, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrInvalidAddress, "fee payer address: %s is not correct", fromAddress)
	}

	if f.ak.GetAccount(ctx, feePayer) == nil {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", feePayer)
	}

	receiverStakeBalances := make(map[string]sdk.Coins)
	for _, receiver := range msg.Outputs {
		noFee := false
		// is fee payer is excluded from fee
		if _, found := f.fek.GetAddress(ctx, receiver.Address); found {
			noFee = true
		}

		receiverStakeBalance := sdk.NewCoins(sdk.NewCoin(domain.DenomStake, sdkmath.NewInt(0)))
		if !noFee {
			receiverAccAddr, err := sdk.AccAddressFromBech32(receiver.Address)
			if err != nil {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInvalidAddress,
					"receiver address: %s is not correct",
					receiver.Address,
				)
			}
			receiverStakeBalance = sdk.NewCoins(
				sdk.NewCoin(
					domain.DenomStake,
					sdkmath.NewInt(f.stakeKeeper.GetFreeStake(ctx, receiverAccAddr).Int64()),
				),
			)
		}

		receiverStakeBalances[receiver.Address] = receiverStakeBalance
	}

	// checking every coin for the fee
	// and getting a total amount of sent denoms
	txFullAmount := sdk.NewCoins()
	feesMap := make(map[string]feesForCoins)
	feesEnabled := make(map[string]struct{}) // used for stats
	for _, output := range msg.Outputs {
		// check every coin for the fee
		for _, coinToSend := range output.Coins {
			if coinToSend.Amount.IsZero() || coinToSend.Amount.IsNegative() {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInvalidCoins,
					"amount have to be > 0, got %s for %s",
					coinToSend.Amount.String(),
					output.Address,
				)
			}

			// will check if sender have enough funds later
			txFullAmount = txFullAmount.Add(coinToSend)

			// is fee payer is excluded from fee
			if _, found := f.fek.GetAddress(ctx, output.Address); found {
				// if output.Address is excluded from fee - set fee to 0
				continue
			}

			if _, found := feesMap[coinToSend.Denom]; !found {
				feesMap[coinToSend.Denom] = make(feesForCoins)
			}

			feeInfo, found := feesMap[coinToSend.Denom][coinToSend.Amount]
			if found {
				continue
			}

			recipient, err := sdk.AccAddressFromBech32(output.Address)
			if err != nil {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInvalidAddress,
					"recipient address: %s is not correct",
					recipient,
				)
			}

			// get receiver stake balance
			recStakeBal, ok := receiverStakeBalances[recipient.String()]
			if !ok {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"receiver stake balance not found for %s",
					recipient.String(),
				)
			}
			// trying to get fee info (percents, min amount) for the coin
			feeInfo.fee, feeInfo.minRefBalance, found = f.fek.GetFees(
				ctx,
				recStakeBal,
				coinToSend,
			)
			if !found {
				continue
			}

			feeAmountFloat, err := strconv.ParseFloat(feeInfo.fee.Fee, 64)
			if err != nil {
				return ctx, sdkioerrors.Wrapf(sdkerrors.ErrLogic, "cant parse fee amount %s", feeInfo.fee.Fee)
			}

			feeAmount := uint64(math.Round(float64(coinToSend.Amount.Uint64()) * feeAmountFloat))
			feeInfo.feeCoin = sdk.NewCoin(coinToSend.Denom, sdkmath.NewIntFromUint64(feeAmount))

			feesMap[coinToSend.Denom][coinToSend.Amount] = feeInfo

			if _, found := feesEnabled[coinToSend.Denom]; !found {
				feesEnabled[coinToSend.Denom] = struct{}{}
			}

			// check every coins amount for the minimum amount if required
			if coinToSend.Amount.LT(sdkmath.NewIntFromUint64(feeInfo.fee.MinAmount)) {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"amount have to be > %d, got %s",
					feeInfo.fee.MinAmount,
					coinToSend.Amount.String(),
				)
			}

		}
	}

	spendableCoins := f.bk.SpendableCoins(ctx, feePayer)
	// check that sender have enough funds for each denom
	for _, denom := range txFullAmount.Denoms() {
		// check if sender have enough funds
		spendable := spendableCoins.AmountOf(denom)
		if spendable.LT(txFullAmount.AmountOf(denom)) {
			return ctx, sdkioerrors.Wrapf(
				sdkerrors.ErrInsufficientFunds,
				"sender have not enough funds, want %s, got %s",
				txFullAmount.AmountOf(denom),
				spendable.String(),
			)
		}

		if denom == domain.DenomStake {
			stakeSpendable := f.stakeKeeper.GetFreeStake(ctx, feePayer)
			if txFullAmount.AmountOf(denom).LT(stakeSpendable) {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"sender have not enough funds, want %s, got %s",
					txFullAmount.AmountOf(denom),
					stakeSpendable.String(),
				)
			}
		}
	}

	// making a copy of the message for deducting fees from the tx
	msgCopy := msg

	fullAmounts := make(map[string]sdk.Coins)
	amountsFeesToSend := make(map[string][]sendFees)

	for _, output := range msgCopy.Outputs {
		// check recipient address and add to the stats if needed
		recipient, err := sdk.AccAddressFromBech32(output.Address)
		if err != nil {
			return ctx, sdkioerrors.Wrapf(
				sdkerrors.ErrInvalidAddress,
				"recipient address: %s is not correct",
				output.Address,
			)
		}
		if !f.ak.HasAccount(ctx, recipient) {
			f.uk.AddNewUserToStat(ctx)
			f.uk.IncrementTotalUsers(ctx)
		}

		// check every coin for the fee
		for _, sendCoin := range output.Coins {
			fullAmounts[output.Address] = fullAmounts[output.Address].Add(sendCoin)

			if feeByDenom, found := feesMap[sendCoin.Denom]; found {
				feeInfo, found := feeByDenom[sendCoin.Amount]
				if !found {
					continue
				}

				msgCopy.Inputs[0].Coins = msgCopy.Inputs[0].Coins.Sub(feeInfo.feeCoin)
				msgCopy.Outputs[i].Coins = msgCopy.Outputs[i].Coins.Sub(feeInfo.feeCoin)
				amountsFeesToSend[output.Address] = append(amountsFeesToSend[output.Address], feeInfo)
			}
		}
	}

	// override amounts
	tx.GetMsgs()[i] = msgCopy

	// set noFee denoms to the stats
	// for _, denom := range txFullAmount.Denoms() {
	//	// is denom is excluded from fee
	//	_, found := feesEnabled[denom]
	//	if !found {
	//		// todo Stats related
	//		fullAmount := sdk.NewCoin(denom, txFullAmount.AmountOf(denom))
	//		f.corek.SetStatsNoFee(ctx, sdk.NewCoins(fullAmount))
	//		continue
	//	}
	// }

	for addressTo, feeCoinsInfo := range amountsFeesToSend {
		for _, feeInfo := range feeCoinsInfo {
			fullAmount := feeInfo.feeCoin.Amount
			fullAmountCoins, found := fullAmounts[addressTo]
			if found {
				fullAmount = fullAmountCoins.AmountOf(feeInfo.feeCoin.Denom)
			}

			ctx, err = f.deductFees(
				ctx,
				feeInfo,
				feePayer,
				addressTo,
				sdk.NewCoins(sdk.NewCoin(feeInfo.feeCoin.Denom, fullAmount)),
			)
			if err != nil {
				return ctx, err
			}
		}
	}

	f.uk.CountUsers(ctx, []string{feePayer.String()})

	return ctx, nil
}
