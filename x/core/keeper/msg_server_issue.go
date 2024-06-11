/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

func (m msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address: %s is not correct", address)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if amount.IsNil() || !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}

	// check if it is a new user // TODO: add stats
	if !m.accountKeeper.HasAccount(ctx, address) {
		m.userKeeper.AddNewUserToStat(ctx)
		m.userKeeper.IncrementTotalUsers(ctx)
	}

	if msg.Denom == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom is empty")
	}

	coinAmount := sdk.NewCoin(msg.Denom, amount)
	coins := sdk.NewCoins(coinAmount)

	// if coinAmount.Denom == types.DenomSTSTWART { // TODO: add stake denom
	//	return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "cannot issue stake")
	// }

	if err = m.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, err
	}

	if err = m.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, coins); err != nil {
		return nil, err
	}

	m.AddIssuedToDailyStats(ctx, coins...)

	if err = ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgIssueResponse{}, nil
}
