package keeper

import (
	"context"

	"cosmossdk.io/log"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) (params types.Params)
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	UserKeeper interface {
		GetAllUser(ctx context.Context) (list []types.User)
		UserAll(ctx context.Context, req *types.QueryAllUserRequest) (*types.QueryAllUserResponse, error)
		User(ctx context.Context, req *types.QueryGetUserRequest) (*types.QueryGetUserResponse, error)
		SetUser(ctx context.Context, user types.User)
		GetUser(ctx context.Context, accountAddress string) (val types.User, found bool)
		RemoveUser(ctx context.Context, accountAddress string)
	}

	ReferralKeeper interface {
		BaseKeeper
		UserKeeper

		SetReferrer(goCtx context.Context, msg *types.MsgSetReferrer) error
	}
)
