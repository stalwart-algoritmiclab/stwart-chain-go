/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
)

// Burn burns coins from the void module account.
func (k Keeper) Burn(ctx sdk.Context, address sdk.AccAddress, amount sdk.Coin) error {
	if amount.IsZero() {
		return nil
	}

	if amount.IsNegative() {
		return fmt.Errorf("invalid amount: %s", amount)
	}

	if amount.Denom == domain.DenomStake {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "cannot burn stake")
	}

	coins := sdk.NewCoins(amount)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, coins); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return err
	}

	k.AddBurnedToDailyStats(ctx, coins...)

	msg := &types.MsgBurn{
		Creator: types.ModuleName,
		Amount:  amount.Amount.Uint64(),
		Denom:   amount.Denom,
		Address: address.String(),
	}
	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return err
	}

	return nil
}

// BurnCoinsWithoutStats burns coins from the voqid module account without updating the daily stats.
func (k Keeper) BurnCoinsWithoutStats(ctx sdk.Context, address sdk.AccAddress, amount sdk.Coin) error {
	if amount.IsZero() {
		return nil
	}

	if amount.IsNegative() {
		return fmt.Errorf("invalid amount: %s", amount)
	}

	if amount.Denom == domain.DenomStake {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cannot burn stake")
	}

	coins := sdk.NewCoins(amount)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, coins); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return err
	}

	msg := &types.MsgBurn{
		Creator: types.ModuleName,
		Amount:  amount.Amount.Uint64(),
		Denom:   amount.Denom,
		Address: address.String(),
	}
	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return err
	}

	return nil
}
