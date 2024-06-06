/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package faucet

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/faucet"
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
					RpcMethod: "TokensAll",
					Use:       "list-tokens",
					Short:     "List all tokens",
				},
				{
					RpcMethod:      "Tokens",
					Use:            "show-tokens [id]",
					Short:          "Shows a tokens by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
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
					RpcMethod:      "CreateTokens",
					Use:            "create-tokens [denom] [amount]",
					Short:          "Create tokens",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "UpdateTokens",
					Use:            "update-tokens [id] [denom] [amount]",
					Short:          "Update tokens",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "denom"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "DeleteTokens",
					Use:            "delete-tokens [id]",
					Short:          "Delete tokens",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "Issue",
					Use:            "issue [address]",
					Short:          "Send a issue tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
