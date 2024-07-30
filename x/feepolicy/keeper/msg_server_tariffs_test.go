/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTariffsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.FeepolicyKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateTariffs{
			Tariffs: &types.Tariff{
				Denom:         domain.DenomStake,
				Id:            1,
				Amount:        "100",
				MinRefBalance: "100",
				Fees: []*types.Fees{
					{
						AmountFrom:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
						Fee:         "",
						RefReward:   "100",
						StakeReward: "100",
						MinAmount:   100,
						NoRefReward: true,
						Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
						Id:          1,
					},
				},
			},
			Creator: creator,
			Denom:   strconv.Itoa(i),
		}
		_, err := srv.CreateTariffs(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetTariffs(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestTariffsMsgServerUpdate(t *testing.T) {
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"

	tests := []struct {
		desc        string
		request     *types.MsgUpdateTariffs
		checkResult bool
		want        types.Tariffs
		err         error
	}{
		{
			desc:        "Completed",
			checkResult: false,
			request: &types.MsgUpdateTariffs{Creator: creator,
				Denom: domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Amount:        "100",
					MinRefBalance: "100",
					Id:            0,
					Fees: []*types.Fees{
						{
							AmountFrom:  "0",
							Fee:         "0.1",
							RefReward:   "0.25",
							StakeReward: "0.25",
							MinAmount:   100,
							NoRefReward: true,
							Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							Id:          0,
						},
						{
							AmountFrom:  "100000000",
							Fee:         "0.01",
							RefReward:   "0.25",
							StakeReward: "0.25",
							MinAmount:   100,
							NoRefReward: true,
							Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							Id:          1,
						},
						{
							AmountFrom:  "10000000000",
							Fee:         "0.001",
							RefReward:   "0.25",
							StakeReward: "0.25",
							MinAmount:   100,
							NoRefReward: true,
							Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							Id:          2,
						},
					},
				},
			},
		},
		{
			desc:        "Completed with id",
			checkResult: true,
			want: types.Tariffs{
				Denom:   domain.DenomStake,
				Creator: creator,
				Tariffs: []*types.Tariff{
					{
						Denom:         domain.DenomStake,
						Id:            0,
						Amount:        "100000",
						MinRefBalance: "100000",
						Fees:          prepareUpdatedFees1(),
					},
					{
						Denom:         domain.DenomStake,
						Id:            1,
						Amount:        "100",
						MinRefBalance: "100",
						Fees:          prepareFees2(),
					},
				},
			},
			request: &types.MsgUpdateTariffs{
				Denom: domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Id:            0,
					Amount:        "100000",
					MinRefBalance: "100000",
					Fees:          prepareUpdatedFees1(),
				},
				Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			},
		},
		{
			desc:        "Completed with id 2",
			checkResult: true,
			want: types.Tariffs{
				Denom:   domain.DenomStake,
				Creator: creator,
				Tariffs: []*types.Tariff{
					{
						Denom:         domain.DenomStake,
						Id:            0,
						Amount:        "100000",
						MinRefBalance: "100000",
						Fees:          prepareFees1(),
					},
					{
						Denom:         domain.DenomStake,
						Id:            1,
						Amount:        "100",
						MinRefBalance: "100",
						Fees:          prepareUpdatedFees2(),
					},
				},
			},
			request: &types.MsgUpdateTariffs{
				Denom: domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Id:            1,
					Amount:        "100",
					MinRefBalance: "100",
					Fees:          prepareUpdatedFees2(),
				},
				Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			},
		},
		{
			desc:        "Completed with extra fee",
			checkResult: true,
			want: types.Tariffs{
				Denom:   domain.DenomStake,
				Creator: creator,
				Tariffs: []*types.Tariff{
					{
						Denom:         domain.DenomStake,
						Id:            0,
						Amount:        "100000",
						MinRefBalance: "100000",
						Fees:          prepareFees1(),
					},
					{
						Denom:         domain.DenomStake,
						Id:            1,
						Amount:        "100",
						MinRefBalance: "100",
						Fees:          prepareUpdatedFees2WithExtraFee(),
					},
				},
			},
			request: &types.MsgUpdateTariffs{
				Denom: domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Id:            1,
					Amount:        "100",
					MinRefBalance: "100",
					Fees:          prepareUpdatedFees2WithExtraFee(),
				},
				Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateTariffs{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateTariffs{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FeepolicyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateTariffs{
				Creator: creator,
				Denom:   domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Amount:        "100000",
					MinRefBalance: "100000",
					Fees:          prepareFees1(),
				}}
			expected2 := &types.MsgCreateTariffs{
				Creator: creator,
				Denom:   domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Amount:        "100",
					MinRefBalance: "100",
					Fees:          prepareFees2(),
				}}
			_, err := srv.CreateTariffs(ctx, expected)
			require.NoError(t, err)

			_, err = srv.CreateTariffs(ctx, expected2)
			require.NoError(t, err)

			_, err = srv.UpdateTariffs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTariffs(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}

			if tc.checkResult {
				tariffs, found := k.GetTariffs(ctx, domain.DenomStake)
				if !found {
					t.Error("get tariffs error")
				}

				t.Log(tariffs)

				assert.DeepEqual(t, tc.want, tariffs)
			}
		})
	}
}

func TestTariffsMsgServerDelete(t *testing.T) {
	creator := "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"

	tests := []struct {
		desc    string
		request *types.MsgDeleteTariffs
		err     error
	}{
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteTariffs{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteTariffs{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.FeepolicyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateTariffs(ctx, &types.MsgCreateTariffs{
				Creator: creator,
				Denom:   domain.DenomStake,
				Tariffs: &types.Tariff{
					Denom:         domain.DenomStake,
					Id:            1,
					Amount:        "100",
					MinRefBalance: "100",
					Fees: []*types.Fees{
						{
							AmountFrom:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							Fee:         "100",
							RefReward:   "100",
							StakeReward: "100",
							MinAmount:   100,
							NoRefReward: true,
							Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
							Id:          1,
						}}},
			})
			require.NoError(t, err)
			_, err = srv.DeleteTariffs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTariffs(ctx,
					tc.request.Denom,
				)
				require.False(t, found)
			}
		})
	}
}

func Test_msgServer_DeleteTariffs(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgDeleteTariffs
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgDeleteTariffsResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgDeleteTariffs{
					Creator:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:    "ccs",
					TariffID: "5",
					FeeID:    "5",
				},
			},
			want:    &types.MsgDeleteTariffsResponse{},
			wantErr: false,
		},
		{
			name: "[FAILED]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgDeleteTariffs{
					Creator:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:    "ccs",
					TariffID: "0",
					FeeID:    "5",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[FAILED]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgDeleteTariffs{
					Creator:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:    "ccs",
					TariffID: "5",
					FeeID:    "0",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgDeleteTariffs{
					Creator:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:    "ccss",
					TariffID: "6",
				},
			},
			want:    &types.MsgDeleteTariffsResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.DeleteTariffs(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTariffs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTariffs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_msgServer_CreateTariffs(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	type args struct {
		goCtx context.Context
		msg   *types.MsgCreateTariffs
	}
	tests := []struct {
		name    string
		args    args
		want    *types.MsgCreateTariffsResponse
		wantErr bool
	}{
		{
			name: "[SUCCESS]",
			args: args{
				goCtx: wctx,
				msg: &types.MsgCreateTariffs{
					Tariffs: &types.Tariff{
						Denom:         domain.DenomStake,
						Id:            1,
						Amount:        "1000000",
						MinRefBalance: "1000000",
						Fees: []*types.Fees{
							{
								AmountFrom:  "1000000",
								Fee:         "",
								RefReward:   "0.25",
								StakeReward: "0.5",
								MinAmount:   150,
								NoRefReward: true,
								Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
								Id:          1,
							},
						},
					},
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   domain.DenomStake,
				},
			},
			want:    &types.MsgCreateTariffsResponse{},
			wantErr: false,
		},
		{
			name: "[SUCCESS] ref reward and stake reward = 1",
			args: args{
				goCtx: wctx,
				msg: &types.MsgCreateTariffs{
					Tariffs: &types.Tariff{
						Denom:         domain.DenomStake,
						Id:            1,
						Amount:        "1000000",
						MinRefBalance: "1000000",
						Fees: []*types.Fees{
							{
								AmountFrom:  "1000000",
								Fee:         "",
								RefReward:   "0.5",
								StakeReward: "0.5",
								MinAmount:   150,
								NoRefReward: true,
								Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
								Id:          1,
							},
						},
					},
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   domain.DenomStake,
				},
			},
			want:    &types.MsgCreateTariffsResponse{},
			wantErr: false,
		},
		{
			name: "[FAILED] stake reward > 1",
			args: args{
				goCtx: wctx,
				msg: &types.MsgCreateTariffs{
					Tariffs: &types.Tariff{
						Denom:         domain.DenomStake,
						Id:            1,
						Amount:        "100000",
						MinRefBalance: "100000",
						Fees: []*types.Fees{
							{
								AmountFrom:  "100000",
								Fee:         "",
								RefReward:   "0.5",
								StakeReward: "0.75",
								MinAmount:   150,
								NoRefReward: true,
								Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
								Id:          1,
							},
						},
					},
					Creator: "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
					Denom:   domain.DenomStake,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateTariffs(tt.args.goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTariffs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTariffs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func prepareFees1() []*types.Fees {
	return []*types.Fees{
		{
			AmountFrom:  "100",
			Fee:         "",
			RefReward:   "0.5",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          0,
		},
		{
			AmountFrom:  "200",
			Fee:         "10",
			RefReward:   "0.5",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          1,
		},
		{
			AmountFrom:  "300",
			Fee:         "10",
			RefReward:   "0.5",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          2,
		},
	}
}

func prepareUpdatedFees1() []*types.Fees {
	return []*types.Fees{
		{
			AmountFrom:  "100",
			Fee:         "",
			RefReward:   "0.3",
			StakeReward: "0.4",
			MinAmount:   101,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          0,
		},
		{
			AmountFrom:  "200",
			Fee:         "10",
			RefReward:   "0.3",
			StakeReward: "0.2",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          1,
		},
		{
			AmountFrom:  "300",
			Fee:         "10",
			RefReward:   "0.1",
			StakeReward: "0.7",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          2,
		},
	}
}

func prepareFees2() []*types.Fees {
	return []*types.Fees{
		{
			AmountFrom:  "0",
			Fee:         "0.1",
			RefReward:   "0.25",
			StakeReward: "0.25",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          0,
		},
		{
			AmountFrom:  "100000000",
			Fee:         "0.01",
			RefReward:   "0.25",
			StakeReward: "0.25",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          1,
		},
		{
			AmountFrom:  "10000000000",
			Fee:         "0.001",
			RefReward:   "0.25",
			StakeReward: "0.25",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          2,
		},
	}
}

func prepareUpdatedFees2() []*types.Fees {
	return []*types.Fees{
		{
			AmountFrom:  "0",
			Fee:         "0.1",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          0,
		},
		{
			AmountFrom:  "100000000",
			Fee:         "0.01",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          1,
		},
		{
			AmountFrom:  "10000000000",
			Fee:         "0.001",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          2,
		},
	}
}

func prepareUpdatedFees2WithExtraFee() []*types.Fees {
	return []*types.Fees{
		{
			AmountFrom:  "0",
			Fee:         "0.1",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          0,
		},
		{
			AmountFrom:  "100000000",
			Fee:         "0.01",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          1,
		},
		{
			AmountFrom:  "10000000000",
			Fee:         "0.001",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          2,
		},
		{
			AmountFrom:  "1000000000000000",
			Fee:         "0.0001",
			RefReward:   "0.37",
			StakeReward: "0.5",
			MinAmount:   100,
			NoRefReward: true,
			Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
			Id:          3,
		},
	}
}
