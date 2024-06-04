package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetReferrer set referrer field to the message creator.
func (k msgServer) SetReferrer(goCtx context.Context, msg *types.MsgSetReferrer) (*types.MsgSetReferrerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.SetReferrer(ctx, msg); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(msg); err != nil {
		return nil, err
	}

	return &types.MsgSetReferrerResponse{}, nil
}
