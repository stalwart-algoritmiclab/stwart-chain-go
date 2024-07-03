/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package rates

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/stalwart-algoritmiclab/stwart-chain-go/api/stwartchain/rates"
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
					RpcMethod: "AddressesAll",
					Use:       "list-addresses",
					Short:     "List all addresses",
				},
				{
					RpcMethod:      "Addresses",
					Use:            "show-addresses [id]",
					Short:          "Shows a addresses by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "RatesAll",
					Use:       "list-rates",
					Short:     "List all rates",
				},
				{
					RpcMethod:      "Rates",
					Use:            "show-rates [id]",
					Short:          "Shows a rates",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}},
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
					RpcMethod:      "CreateAddresses",
					Use:            "create-addresses [address]",
					Short:          "Create addresses",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateAddresses",
					Use:            "update-addresses [id] [address]",
					Short:          "Update addresses",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "DeleteAddresses",
					Use:            "delete-addresses [id]",
					Short:          "Delete addresses",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateRates",
					Use:            "create-rates [denom] [rate] [decimals]",
					Short:          "Create a new rates",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "rate"}, {ProtoField: "decimals"}},
				},
				{
					RpcMethod:      "UpdateRates",
					Use:            "update-rates [denom] [rate] [decimals]",
					Short:          "Update rates",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "rate"}, {ProtoField: "decimals"}},
				},
				{
					RpcMethod:      "DeleteRates",
					Use:            "delete-rates [denom]",
					Short:          "Delete rates",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
