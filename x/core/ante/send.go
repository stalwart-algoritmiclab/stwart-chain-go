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

func (f CoreDecorator) processMsgSend(ctx sdk.Context, tx sdk.Tx, msgSend sdk.Msg, i int) (sdk.Context, error) {
	msg, ok := msgSend.(*types.MsgSend)
	if !ok {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unexpected MsgSend type %s", msg.FromAddress)
	}
	// check recipient address and add to the stats if needed
	recipient, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrInvalidAddress, "recipient address: %s is not correct", msg.FromAddress)
	}
	if !f.ak.HasAccount(ctx, recipient) {
		f.uk.AddNewUserToStat(ctx)
		f.uk.IncrementTotalUsers(ctx)
	}

	feePayer, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrInvalidAddress, "fee payer address: %s is not correct", msg.FromAddress)
	}

	if f.ak.GetAccount(ctx, feePayer) == nil {
		return ctx, sdkioerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", feePayer)
	}

	noFee := false
	// is fee payer is excluded from fee
	if _, found := f.fek.GetAddress(ctx, msg.ToAddress); found {
		noFee = true
	}

	receiverStakeBalance := sdk.NewCoins(
		sdk.NewCoin(
			domain.DenomStake,
			sdkmath.NewInt(f.stakeKeeper.GetFreeStake(ctx, recipient).Int64()),
		),
	)

	txFullAmount := msg.Amount

	// check every coin for the fee
	for _, sendCoin := range msg.Amount {
		tokens := sendCoin.Amount

		if tokens.IsZero() || tokens.IsNegative() {
			return ctx, sdkioerrors.Wrapf(
				sdkerrors.ErrInvalidCoins,
				"amount have to be > 0, got %s",
				tokens.String(),
			)
		}

		assets := f.bk.SpendableCoins(ctx, feePayer)
		spendable := assets.AmountOf(sendCoin.Denom)
		if spendable.LT(tokens) {
			return ctx, sdkioerrors.Wrapf(
				sdkerrors.ErrInsufficientFunds,
				"sender have not enough funds, want %s, got %s",
				tokens.String(),
				spendable.String(),
			)
		}

		if sendCoin.Denom == domain.DenomStake {
			spendable = f.stakeKeeper.GetFreeStake(ctx, feePayer)
			if spendable.LT(tokens) {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"sender have not enough funds, want %s, got %s",
					tokens.String(),
					spendable.String(),
				)
			}
		}

		feeCoin := sdk.NewInt64Coin(sendCoin.Denom, 0)

		// get fee percent
		fee, minRefBalance, found := f.fek.GetFees(ctx, receiverStakeBalance, sendCoin)
		if noFee {
			// if fee payer is excluded from fee - set fee to 0
			found = false
		}
		if found {
			if tokens.LT(sdkmath.NewIntFromUint64(fee.MinAmount)) {
				return ctx, sdkioerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"amount have to be > %d, got %s",
					fee.MinAmount,
					tokens.String(),
				)
			}
			//
			feeAmountFloat, err := strconv.ParseFloat(fee.Fee, 64)
			if err != nil {
				return ctx, sdkioerrors.Wrapf(sdkerrors.ErrLogic, "cant parse fee amount %s", fee.Fee)
			}

			feeAmount := uint64(math.Round(float64(tokens.Uint64()) * feeAmountFloat))
			feeCoin.Amount = sdkmath.NewIntFromUint64(feeAmount)

			msg.Amount = msg.Amount.Sub(feeCoin)
			tx.GetMsgs()[i] = msg

			feeInfo := sendFees{feeCoin: feeCoin, minRefBalance: minRefBalance, fee: fee}
			ctx, err := f.deductFees(ctx, feeInfo, feePayer, msg.ToAddress, txFullAmount)
			if err != nil {
				return ctx, err
			}
		}
		// else {
		//	// todo stats
		//	f.corek.SetStatsNoFee(ctx, txFullAmount)
		// }
	}

	f.uk.CountUsers(ctx, []string{feePayer.String()})
	return ctx, nil
}
