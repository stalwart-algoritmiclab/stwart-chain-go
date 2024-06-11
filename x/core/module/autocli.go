/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package core

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/core"
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
					RpcMethod:      "ModulesAddresses",
					Use:            "modules-addresses",
					Short:          "Query modulesAddresses",
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
				{
					RpcMethod:      "Issue",
					Use:            "issue [amount] [denom] [address]",
					Short:          "Send a Issue tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "denom"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "Withdraw",
					Use:            "withdraw [amount] [denom] [address]",
					Short:          "Send a Withdraw tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "denom"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "Send",
					Use:            "send [from] [to] [amount] [denom]",
					Short:          "Send a Send tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "from"}, {ProtoField: "to"}, {ProtoField: "amount"}, {ProtoField: "denom"}},
				},
				{
					RpcMethod:      "Refund",
					Use:            "refund [from] [to] [amount]",
					Short:          "Send a Refund tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "from"}, {ProtoField: "to"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "Fees",
					Use:            "fees [comission] [address-to]",
					Short:          "Send a Fees tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "comission"}, {ProtoField: "addressTo"}},
				},
				{
					RpcMethod:      "RefReward",
					Use:            "ref-reward [amount] [referrer]",
					Short:          "Send a RefReward tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "referrer"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
