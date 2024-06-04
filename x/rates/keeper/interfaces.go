package keeper

import (
	"context"

	"cosmossdk.io/log"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

type (
	BaseKeeper interface {
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	RatesKeeper interface {
		BaseKeeper

		CreateRates(goCtx context.Context, msg *types.MsgCreateRates) (*types.MsgCreateRatesResponse, error)
		UpdateRates(goCtx context.Context, msg *types.MsgUpdateRates) (*types.MsgUpdateRatesResponse, error)
		DeleteRates(goCtx context.Context, msg *types.MsgDeleteRates) (*types.MsgDeleteRatesResponse, error)
		RatesAll(ctx context.Context, req *types.QueryAllRatesRequest) (*types.QueryAllRatesResponse, error)
		Rates(ctx context.Context, req *types.QueryGetRatesRequest) (*types.QueryGetRatesResponse, error)
	}
)
