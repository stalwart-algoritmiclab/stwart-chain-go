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

func (m msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "creator: %s is not allowed", msg.Creator)
	}

	addressTo, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address to: %s is not correct", addressTo)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if amount.IsNil() || !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", msg.Amount)
	}

	if amount.IsZero() {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount is zero")
	}

	if msg.Denom == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom is empty")
	}

	coins := sdk.NewCoins(sdk.NewCoin(msg.Denom, amount))

	var moduleName string
	for _, info := range m.modulesList {
		if msg.From == info.Address {
			moduleName = info.Name
			break
		}
	}

	if moduleName == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address: %s", msg.From)
	}

	if err = m.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, addressTo, coins); err != nil {
		return nil, err
	}

	if err = ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgSendResponse{}, nil
}
