/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package systemrewards

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/stalwart-algoritmiclab/stwart-chain-go/api/stwartchain/systemrewards"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "StatsAll",
					Use:       "list-stats",
					Short:     "List all stats",
				},
				{
					RpcMethod:      "Stats",
					Use:            "show-stats [date]",
					Short:          "Shows a stats",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},
				{
					RpcMethod:      "StatsByDate",
					Use:            "stats-by-date [start-date] [end-date]",
					Short:          "get daily statistics between start-date and end-date",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "startDate"}, {ProtoField: "endDate"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
