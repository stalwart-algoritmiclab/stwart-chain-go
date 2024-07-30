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

var _ sdk.Msg = &MsgCreatePollsParams{}

func NewMsgCreatePollsParams(creator string, proposerDeposit []sdk.Coin, burnVeto bool) *MsgCreatePollsParams {
	return &MsgCreatePollsParams{
		Creator:         creator,
		ProposerDeposit: proposerDeposit,
		BurnVeto:        burnVeto,
	}
}

func (msg *MsgCreatePollsParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePollsParams{}

func NewMsgUpdatePollsParams(creator string, proposerDeposit []sdk.Coin, burnVeto bool) *MsgUpdatePollsParams {
	return &MsgUpdatePollsParams{
		Creator:         creator,
		ProposerDeposit: proposerDeposit,
		BurnVeto:        burnVeto,
	}
}

func (msg *MsgUpdatePollsParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePollsParams{}

func NewMsgDeletePollsParams(creator string) *MsgDeletePollsParams {
	return &MsgDeletePollsParams{
		Creator: creator,
	}
}

func (msg *MsgDeletePollsParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
