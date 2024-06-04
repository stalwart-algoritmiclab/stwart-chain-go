package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// get list of tokens
	tokens := k.Keeper.GetAllTokens(ctx)

	for _, token := range tokens {
		address, _ := sdk.AccAddressFromBech32(msg.Address)

		amount, ok := math.NewIntFromString(token.Amount)
		if !ok {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", token.Amount)
		}

		coinAmount := sdk.NewCoin(token.Denom, amount)
		coins := sdk.NewCoins(coinAmount)

		err := k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
		if err != nil {
			return nil, err
		}

		err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, coins)
		if err != nil {
			return nil, err
		}

	}

	return &types.MsgIssueResponse{}, nil
}
