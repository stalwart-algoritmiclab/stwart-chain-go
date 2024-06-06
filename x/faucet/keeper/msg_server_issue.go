/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

func (m msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// get list of tokens
	tokens := m.Keeper.GetAllTokens(ctx)

	for _, token := range tokens {
		address, err := sdk.AccAddressFromBech32(msg.Address)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address: %s", err.Error())
		}

		amount, ok := math.NewIntFromString(token.Amount)
		if !ok || amount.IsZero() {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", token.Amount)
		}

		coins := sdk.NewCoins(sdk.NewCoin(token.Denom, amount))

		if err = m.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
			return nil, err
		}

		if err = m.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, coins); err != nil {
			return nil, err
		}

	}

	return &types.MsgIssueResponse{}, nil
}
