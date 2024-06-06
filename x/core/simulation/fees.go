/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package simulation

import (
	"math/rand"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgFees(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		//simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgFees{
			Comission: sdk.Coin{},
		}

		// TODO: Handling the Fees simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "Fees simulation not implemented"), nil, nil
	}
}
