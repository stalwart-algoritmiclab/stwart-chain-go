/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoll{}

func NewMsgCreatePoll(creator string, title string, description string, votingStartTime string, votingPeriod string, minVoteAmount uint64, minAdressesCount uint64, minVoteCoinsAmount uint64, options []Options) *MsgCreatePoll {
	return &MsgCreatePoll{
		Creator:            creator,
		Title:              title,
		Description:        description,
		VotingStartTime:    votingStartTime,
		VotingPeriod:       votingPeriod,
		MinVoteAmount:      minVoteAmount,
		MinAdressesCount:   minAdressesCount,
		MinVoteCoinsAmount: minVoteCoinsAmount,
		Options:            options,
	}
}

func (msg *MsgCreatePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
