/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package polls

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/stalwart-algoritmiclab/stwart-chain-go/api/stwartchain/polls"
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
					RpcMethod: "PollsParams",
					Use:       "show-polls-params",
					Short:     "show pollsParams",
				},
				{
					RpcMethod: "VotesAll",
					Use:       "list-votes",
					Short:     "List all votes",
				},
				{
					RpcMethod:      "Votes",
					Use:            "show-votes [id]",
					Short:          "Shows a votes by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "OptionsAll",
					Use:       "list-options",
					Short:     "List all options",
				},
				{
					RpcMethod:      "Options",
					Use:            "show-options [id]",
					Short:          "Shows a options by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "PollsAll",
					Use:       "list-polls",
					Short:     "List all polls",
				},
				{
					RpcMethod:      "Polls",
					Use:            "show-polls [id]",
					Short:          "Shows a polls by id",
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
					RpcMethod:      "CreatePollsParams",
					Use:            "create-polls-params [proposerDeposit] [burnVeto] [minDaysDuration] [maxDaysDuration] [maxDaysPending]",
					Short:          "Create pollsParams",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proposerDeposit"}, {ProtoField: "burnVeto"}, {ProtoField: "minDaysDuration"}, {ProtoField: "maxDaysDuration"}, {ProtoField: "maxDaysPending"}},
				},
				{
					RpcMethod:      "UpdatePollsParams",
					Use:            "update-polls-params [proposerDeposit] [burnVeto] [minDaysDuration] [maxDaysDuration] [maxDaysPending]",
					Short:          "Update pollsParams",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proposerDeposit"}, {ProtoField: "burnVeto"}, {ProtoField: "minDaysDuration"}, {ProtoField: "maxDaysDuration"}, {ProtoField: "maxDaysPending"}},
				},
				{
					RpcMethod: "DeletePollsParams",
					Use:       "delete-polls-params",
					Short:     "Delete pollsParams",
				},
				{
					RpcMethod:      "CreatePoll",
					Use:            "create-poll [title] [description] [voting-start-time] [voting-period] [min-vote-amount] [min-adresses-count] [min-vote-coins-amount] [options]",
					Short:          "Send a createPoll tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "votingStartTime"}, {ProtoField: "votingPeriod"}, {ProtoField: "minVoteAmount"}, {ProtoField: "minAdressesCount"}, {ProtoField: "minVoteCoinsAmount"}, {ProtoField: "options"}},
				},
				{
					RpcMethod:      "Vote",
					Use:            "vote [poll-id] [option-id] [amount]",
					Short:          "Send a Vote tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pollId"}, {ProtoField: "optionId"}, {ProtoField: "amount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
