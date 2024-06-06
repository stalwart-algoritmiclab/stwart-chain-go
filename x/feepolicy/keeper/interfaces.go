/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	FeePolicyKeeper interface {
		BaseKeeper

		Address(goCtx context.Context, req *types.QueryGetAddressRequest) (*types.QueryGetAddressResponse, error)
		AddressesAll(goCtx context.Context, req *types.QueryAllAddressesRequest) (*types.QueryAllAddressesResponse, error)
		AddressByID(goCtx context.Context, req *types.QueryGetAddressByIDRequest) (*types.QueryGetAddressResponse, error)
		AppendAddress(ctx sdk.Context, address types.Address) uint64
		GetAddress(ctx sdk.Context, address string) (val types.Address, found bool)
		GetAddressByID(ctx sdk.Context, id uint64) (val types.Address, found bool)
		GetAddressCount(ctx sdk.Context) uint64
		GetAllAddress(ctx sdk.Context) (list []types.Address)
		RemoveAddress(ctx sdk.Context, id uint64)
		SetAddress(ctx sdk.Context, address types.Address)
		SetAddressCount(ctx sdk.Context, count uint64)
	}
)
