/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

type (
	BaseKeeper interface {
		GetAuthority() string
		GetParams(ctx context.Context) types.Params
		Logger() log.Logger
		Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		SetParams(ctx context.Context, params types.Params) error
	}

	SecuredKeeper interface {
		BaseKeeper

		Addresses(goCtx context.Context, req *types.QueryGetAddressesRequest) (*types.QueryGetAddressesResponse, error)
		AddressesAll(goCtx context.Context, req *types.QueryAllAddressesRequest) (*types.QueryAllAddressesResponse, error)
		AddressesByAddress(goCtx context.Context, req *types.QueryGetAddressRequest) (*types.QueryGetAddressesResponse, error)
		AppendAddresses(ctx context.Context, addresses types.Addresses) uint64
		GetAddresses(ctx context.Context, id uint64) (val types.Addresses, found bool)
		GetAddressesByAddress(ctx sdk.Context, address string) (val types.Addresses, found bool)
		GetAddressesCount(ctx context.Context) uint64
		GetAllAddresses(ctx context.Context) (list []types.Addresses)
		RemoveByID(ctx context.Context, id uint64)
		SetAddresses(ctx context.Context, addresses types.Addresses)
		SetAddressesCount(ctx context.Context, count uint64)
	}
)
