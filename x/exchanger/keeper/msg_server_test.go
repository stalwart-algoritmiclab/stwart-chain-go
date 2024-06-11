/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	math2 "math"
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
	ratesmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.ExchangerKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func TestFormulasIn(t *testing.T) {
	amount := int64(1000)
	amountToExchange := math.NewInt(amount)
	rate := ratesmoduletypes.Rates{
		Denom:    "denom",
		Decimals: 8,
		Rate:     1,
	}
	amountToExchangeFloat := float64(amountToExchange.Uint64()) / math2.Pow10(int(rate.Decimals)) * rate.Rate
	issueAmount := math.NewInt(int64(amountToExchangeFloat * math2.Pow10(domain.DenomStakeDecimals)))

	require.Equal(t, amount, issueAmount.Int64())
}
