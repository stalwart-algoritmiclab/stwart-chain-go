/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	math2 "math"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
	ratestypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
)

func (k msgServer) Exchange(goCtx context.Context, msg *types.MsgExchange) (*types.MsgExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addressTo, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	amountToExchange, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid exchange amount (%s)", msg.Amount)
	}

	switch msg.DenomTo {
	case domain.DenomStableIndex:
		// check if user has denom amount to exchange
		assets := k.bankKeeper.SpendableCoins(ctx, addressTo)
		spendable := assets.AmountOf(msg.Denom)
		if spendable.LT(amountToExchange) {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to exchange")
		}

		ratesInfo, err := k.ratesKeeper.Rates(ctx, &ratestypes.QueryGetRatesRequest{Denom: msg.Denom})
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "rate for %s not found", msg.Denom)
		}

		rate := ratesInfo.Rates
		amountToExchangeFloat := (float64(amountToExchange.Uint64()) / math2.Pow10(int(rate.Decimals))) * rate.Rate
		if amountToExchangeFloat < 0.1 {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "exchange amount less than 0.1")
		}

		issueAmount := math.NewInt(int64(amountToExchangeFloat * math2.Pow10(domain.DenomStableDecimals)))
		if issueAmount.IsZero() {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid exchange amount")
		}

		issueCoin := sdk.NewCoin(domain.DenomStableIndex, issueAmount)
		coins := sdk.NewCoins(issueCoin)

		err = k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
		if err != nil {
			return nil, err
		}

		// send the exchanged amount to the module
		err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, addressTo, types.ModuleName, sdk.NewCoins(sdk.NewCoin(msg.Denom, amountToExchange)))
		if err != nil {
			return nil, err
		}

		// send the exchanged amount to the user
		err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addressTo, coins)
		if err != nil {
			return nil, err
		}

	default:
		// check stake amount to exchange
		assets := k.bankKeeper.SpendableCoins(ctx, addressTo)
		spendable := assets.AmountOf(domain.DenomStableIndex)
		if spendable.LT(amountToExchange) {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to exchange")
		}

		ratesInfo, err := k.ratesKeeper.Rates(ctx, &ratestypes.QueryGetRatesRequest{Denom: msg.DenomTo})
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "rate for %s not found", msg.Denom)
		}

		rate := ratesInfo.Rates
		amountToExchangeFloat := (float64(amountToExchange.Uint64()) / math2.Pow10(domain.DenomStableDecimals))
		if amountToExchangeFloat < 0.1 {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "exchange amount less than 0.1")
		}
		amountToExchangeTo := amountToExchangeFloat / rate.Rate

		transferAmount := math.NewInt(int64(amountToExchangeTo * math2.Pow10(int(rate.Decimals))))
		if transferAmount.IsZero() {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid exchange amount")
		}

		transferCoins := sdk.NewCoins(sdk.NewCoin(msg.DenomTo, transferAmount))

		// check if module has enough coins to transfer
		moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
		moduleAssets := k.bankKeeper.SpendableCoins(ctx, moduleAddress)
		spendable = moduleAssets.AmountOf(msg.DenomTo)
		if spendable.LT(transferAmount) {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to exchange")
		}

		// send tokens to the user
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addressTo, transferCoins)
		if err != nil {
			return nil, err
		}

		coins := sdk.NewCoins(sdk.NewCoin(domain.DenomStableIndex, amountToExchange))

		// get tokens from the user
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addressTo, types.ModuleName, coins)
		if err != nil {
			return nil, err
		}

		// burn the exchanged stake amount
		err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
		if err != nil {
			return nil, err
		}
	}

	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgExchangeResponse{}, nil
}
