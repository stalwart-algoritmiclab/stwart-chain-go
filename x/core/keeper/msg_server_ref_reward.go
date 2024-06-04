package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RefReward(goCtx context.Context, msg *types.MsgRefReward) (*types.MsgRefRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRefRewardResponse{}, nil
}