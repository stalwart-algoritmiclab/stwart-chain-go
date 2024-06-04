package keeper

import (
	"context"

	"cosmossdk.io/log"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

type (
	BaseKeeper interface {
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	FaucetKeeper interface {
		BaseKeeper
		Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error)
		CreateTokens(goCtx context.Context, msg *types.MsgCreateTokens) (*types.MsgCreateTokensResponse, error)
		UpdateTokens(goCtx context.Context, msg *types.MsgUpdateTokens) (*types.MsgUpdateTokensResponse, error)
		DeleteTokens(goCtx context.Context, msg *types.MsgDeleteTokens) (*types.MsgDeleteTokensResponse, error)
	}
)
