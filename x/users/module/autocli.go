package users

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/users"
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
					Short:     "List all Stats",
				},
				{
					RpcMethod:      "Stats",
					Use:            "show-stats [date]",
					Short:          "Shows a Stats",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},
				{
					RpcMethod:      "StatsByDate",
					Use:            "stats-by-date [start-date] [end-date]",
					Short:          "get daily statistics between start-date and end-date",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "startDate"}, {ProtoField: "endDate"}},
				},

				{
					RpcMethod: "UniqueUsersAll",
					Use:       "list-unique-users",
					Short:     "List all UniqueUsers",
				},
				{
					RpcMethod:      "UniqueUsers",
					Use:            "show-unique-users [date]",
					Short:          "Shows a UniqueUsers",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},
				{
					RpcMethod:      "Total",
					Use:            "total",
					Short:          "Query Total",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
			},
		},
	}
}
