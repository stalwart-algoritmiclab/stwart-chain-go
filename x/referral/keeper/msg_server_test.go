package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.ReferralKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}