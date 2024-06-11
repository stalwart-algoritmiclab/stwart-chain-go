/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package ante

import (
	"strconv"

	sdkioerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	coretypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	systemrewardsmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"
)

func (f CoreDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {
	for i, msgSend := range tx.GetMsgs() {
		switch msgSend.(type) {
		case *types.MsgSend:
			return f.processMsgSend(ctx, tx, msgSend, i)
		case *types.MsgMultiSend:
			return f.processMsgMultiSend(ctx, tx, msgSend, i)
		}
	}

	return next(ctx, tx, simulate)
}

// deductFees - deduct fees from the fee payer and send to the system modules or referrers
func (f CoreDecorator) deductFees(
	ctx sdk.Context,
	feeInfo sendFees,
	feePayer sdk.AccAddress,
	addressTo string,
	txFullAmount sdk.Coins,
) (sdk.Context, error) {
	feeCoin := feeInfo.feeCoin
	fees := feeInfo.fee

	if addressTo == "" {
		return ctx, nil
	}
	// todo STATS related
	// f.corek.SetStatsFee(ctx, sdk.NewCoins(feeCoin), txFullAmount)

	// log fee event
	_ = ctx.EventManager().EmitTypedEvents(&coretypes.MsgFees{
		Comission: feeCoin,
		AddressTo: addressTo,
	})

	// get stake reward percent
	stakeReward, _ := strconv.ParseFloat(fees.StakeReward, 64)
	refReward, _ := strconv.ParseFloat(fees.RefReward, 64)
	isNoRefRewards := fees.NoRefReward

	// calculate fee for burn, referral reward and stake reward
	feeAmount := feeCoin.Amount.Uint64()
	feeStakeReward := uint64(float64(feeAmount) * stakeReward)
	feeRefReward := uint64(float64(feeAmount) * refReward) // 1 use referrals =false
	feeVoid := feeAmount - (feeStakeReward + feeRefReward) // 100% - 75% = 25%
	if feeVoid > 0 {
		coinVoid := sdk.NewCoin(feeCoin.Denom, sdkmath.NewIntFromUint64(feeVoid))

		// send fee for void and burn
		if err := f.corek.Burn(ctx, feePayer, coinVoid); err != nil {
			return ctx, err
		}
	}

	// todo STAKE related
	// if feeStakeReward > 0 {
	// 	coinStakeReward := sdk.NewCoins(sdk.NewCoin(feeCoin.Denom, sdkmath.NewIntFromUint64(feeStakeReward)))
	// 	// send fee to the stake reward module
	// 	if err := f.bk.SendCoinsFromAccountToModule(
	// 		ctx,
	// 		feePayer,
	// 		f.stakeRewardsModuleName,
	// 		coinStakeReward,
	// 	); err != nil {
	// 		return ctx, err
	// 	}
	// 	f.stakeRewarsKeeper.AddStats(ctx, coinStakeReward...)
	// }

	if feeRefReward == 0 {
		return ctx, nil
	}

	coinRefReward := sdk.NewCoins(sdk.NewCoin(feeCoin.Denom, sdkmath.NewIntFromUint64(feeRefReward)))

	// send fee to the referrer
	var (
		err error
		ok  = false
	)

	if !isNoRefRewards {
		ok, err = f.sendReferralReward(ctx, feePayer, addressTo, coinRefReward, feeInfo.minRefBalance)
		if err != nil {
			return ctx, err
		}
	}
	if !ok {
		// if no referral - send to the system ref reward module
		if err := f.bk.SendCoinsFromAccountToModule(
			ctx,
			feePayer,
			systemrewardsmoduletypes.ModuleName,
			coinRefReward,
		); err != nil {
			return ctx, err
		}
		f.rewk.AddStats(ctx, coinRefReward...)
	}

	return ctx, nil
}

func (f CoreDecorator) sendReferralReward(
	ctx sdk.Context,
	addressFrom sdk.AccAddress,
	addressTo string,
	refReward sdk.Coins,
	minRefBalance sdk.Coin,
) (bool, error) {
	// SendRefReward send rewards to referrer or module using referral address.
	user, found := f.refk.GetUser(ctx, addressTo)
	if !found {
		return false, nil
	}

	if user.Referrer == "" {
		return false, nil
	}

	if found = f.ak.HasAccount(ctx, addressFrom); !found {
		return false, sdkioerrors.Wrapf(sdkerrors.ErrUnknownAddress, "sender account address: %s", addressFrom)
	}

	referrerAddress, err := sdk.AccAddressFromBech32(user.Referrer)
	if err != nil {
		return false, sdkioerrors.Wrapf(sdkerrors.ErrInvalidAddress, "referrer account address: %s", err.Error())
	}

	// check if referrer have enough minRefBalance
	// if minRefBalance.Amount.GT(sdkmath.NewInt(0)) {
	//	// todo STAKE related
	//	refStake, found := f.stakek.GetStakes(ctx, referrerAddress.String())
	//	if !found {
	//		return false, nil
	//	}
	//
	//	if refStake.Stake.StakeAmount.IsLT(minRefBalance) {
	//		return false, nil
	//	}
	// }

	if err = f.bk.SendCoins(ctx, addressFrom, referrerAddress, refReward); err != nil {
		return false, err
	}

	// log ref reward coins
	for _, coin := range refReward {
		_ = ctx.EventManager().EmitTypedEvents(&coretypes.MsgRefReward{
			Creator:  addressTo,
			Amount:   coin,
			Referrer: referrerAddress.String(),
		})
	}

	return true, nil
}
