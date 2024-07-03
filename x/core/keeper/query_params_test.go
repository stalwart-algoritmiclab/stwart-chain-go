/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)

	_, err = keeper.Params(ctx, nil)
	require.Error(t, err)
}
