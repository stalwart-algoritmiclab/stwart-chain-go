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

var _ sdk.Msg = &MsgCreateRates{}

func NewMsgCreateRates(
	creator string,
	denom string,
	rate string,
	decimals int32,

) *MsgCreateRates {
	return &MsgCreateRates{
		Creator:  creator,
		Denom:    denom,
		Rate:     rate,
		Decimals: decimals,
	}
}

func (msg *MsgCreateRates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRates{}

func NewMsgUpdateRates(
	creator string,
	denom string,
	rate string,

) *MsgUpdateRates {
	return &MsgUpdateRates{
		Creator: creator,
		Denom:   denom,
		Rate:    rate,
	}
}

func (msg *MsgUpdateRates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRates{}

func NewMsgDeleteRates(
	creator string,
	denom string,

) *MsgDeleteRates {
	return &MsgDeleteRates{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgDeleteRates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
