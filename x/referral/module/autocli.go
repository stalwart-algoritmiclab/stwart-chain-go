/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package referral

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/referral"
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
					RpcMethod: "UserAll",
					Use:       "list-user",
					Short:     "List all User",
				},
				{
					RpcMethod:      "User",
					Use:            "show-user [id]",
					Short:          "Shows a User",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "accountAddress"}},
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
					RpcMethod:      "CreateUser",
					Use:            "create-user [accountAddress] [referrer] [referrals]",
					Short:          "Create a new User",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "accountAddress"}, {ProtoField: "referrer"}, {ProtoField: "referrals"}},
				},
				{
					RpcMethod:      "UpdateUser",
					Use:            "update-user [accountAddress] [referrer] [referrals]",
					Short:          "Update User",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "accountAddress"}, {ProtoField: "referrer"}, {ProtoField: "referrals"}},
				},
				{
					RpcMethod:      "DeleteUser",
					Use:            "delete-user [accountAddress]",
					Short:          "Delete User",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "accountAddress"}},
				},
				{
					RpcMethod:      "SetReferrer",
					Use:            "set-referrer [referrer-address] [referral-address]",
					Short:          "Send a SetReferrer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "referrerAddress"}, {ProtoField: "referralAddress"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
