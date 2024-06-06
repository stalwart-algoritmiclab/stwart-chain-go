/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"reflect"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/rand"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

func TestAddressesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SecuredKeeper(t)
	msgs := createNAddresses(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAddressesRequest
		response *types.QueryGetAddressesResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAddressesRequest{Id: msgs[0].Id},
			response: &types.QueryGetAddressesResponse{Addresses: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetAddressesRequest{Id: msgs[1].Id},
			response: &types.QueryGetAddressesResponse{Addresses: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetAddressesRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Addresses(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAddressesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SecuredKeeper(t)
	msgs := createNAddresses(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAddressesRequest {
		return &types.QueryAllAddressesRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AddressesAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Addresses), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Addresses),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AddressesAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Addresses), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Addresses),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AddressesAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Addresses),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AddressesAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestKeeper_AddressesByAddress(t *testing.T) {
	k, _, ctx := setupMsgServer(t)
	account, _ := simtypes.RandomAcc(rand.NewRand(), simtypes.RandomAccounts(rand.NewRand(), 1))
	a := types.Addresses{
		Id:      1,
		Address: []string{account.Address.String()},
		Creator: account.Address.String(),
	}
	k.SetAddresses(ctx, a)

	type args struct {
		req *types.QueryGetAddressRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryGetAddressesResponse
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				req: &types.QueryGetAddressRequest{
					Address: account.Address.String(),
				},
			},
			want: &types.QueryGetAddressesResponse{
				Addresses: types.Addresses{
					Id:      a.Id,
					Address: a.Address,
					Creator: a.Creator,
				},
			},
			wantErr: false,
		},
		{
			name: "Empty request",
			args: args{
				req: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid address",
			args: args{
				req: &types.QueryGetAddressRequest{
					Address: "Some address",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := k.AddressesByAddress(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressesByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressesByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeeper_AddressesAll(t *testing.T) {
	k, _, ctx := setupMsgServer(t)
	accounts := simtypes.RandomAccounts(rand.NewRand(), 3)
	if len(accounts) < 3 {
		t.Errorf("Not enough accounts: %v < 3", len(accounts))
		return
	}

	addresses := make([]string, 0, len(accounts))
	for _, a := range accounts {
		addresses = append(addresses, a.Address.String())
	}

	creator := addresses[0]

	k.SetAddresses(ctx, types.Addresses{
		Id:      1,
		Address: addresses,
		Creator: creator,
	})

	type args struct {
		req *types.QueryAllAddressesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success with limit",
			args: args{
				req: &types.QueryAllAddressesRequest{Pagination: &query.PageRequest{Limit: 2}},
			},
			wantErr: false,
		},
		{
			name: "Success with offset",
			args: args{
				req: &types.QueryAllAddressesRequest{Pagination: &query.PageRequest{Offset: 1}},
			},
			wantErr: false,
		},
		{
			name: "Success with reverse",
			args: args{
				req: &types.QueryAllAddressesRequest{Pagination: &query.PageRequest{Reverse: true}},
			},
			wantErr: false,
		},
		{
			name: "Success with key",
			args: args{
				req: &types.QueryAllAddressesRequest{Pagination: &query.PageRequest{Key: []byte(creator)}},
			},
			wantErr: false,
		},
		{
			name: "Success without pagination",
			args: args{
				req: &types.QueryAllAddressesRequest{},
			},
			wantErr: false,
		},
		{
			name: "Failed with nil request",
			args: args{
				req: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := k.AddressesAll(ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("AddressesAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestKeeper_Addresses(t *testing.T) {
	k, _, ctx := setupMsgServer(t)
	account, _ := simtypes.RandomAcc(rand.NewRand(), simtypes.RandomAccounts(rand.NewRand(), 1))
	k.SetAddresses(ctx, types.Addresses{
		Id:      1,
		Address: []string{account.Address.String()},
		Creator: account.Address.String(),
	})

	type args struct {
		req *types.QueryGetAddressesRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				req: &types.QueryGetAddressesRequest{Id: 1},
			},
			wantErr: false,
		},
		{
			name: "Failed with not found id",
			args: args{
				req: &types.QueryGetAddressesRequest{Id: 100},
			},
			wantErr: true,
		},
		{
			name: "Failed with nil request",
			args: args{
				req: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := k.Addresses(ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Addresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
