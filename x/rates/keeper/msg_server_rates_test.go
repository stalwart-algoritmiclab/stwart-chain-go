package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRatesMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RatesKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRates{Creator: creator,
			Denom: strconv.Itoa(i),
		}
		_, err := srv.CreateRates(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRates(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRatesMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateRates
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRates{Creator: creator,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRates{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRates{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RatesKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateRates{Creator: creator,
				Denom: strconv.Itoa(0),
			}
			_, err := srv.CreateRates(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRates(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRates(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRatesMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteRates
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRates{Creator: creator,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRates{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRates{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RatesKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateRates(ctx, &types.MsgCreateRates{Creator: creator,
				Denom: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteRates(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRates(ctx,
					tc.request.Denom,
				)
				require.False(t, found)
			}
		})
	}
}
