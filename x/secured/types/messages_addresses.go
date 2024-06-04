package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAddresses{}

func NewMsgCreateAddresses(creator string, address []string) *MsgCreateAddresses {
	return &MsgCreateAddresses{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateAddresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAddresses{}

func NewMsgUpdateAddresses(creator string, id uint64, address []string) *MsgUpdateAddresses {
	return &MsgUpdateAddresses{
		Id:      id,
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateAddresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAddresses{}

func NewMsgDeleteAddresses(creator string, id uint64) *MsgDeleteAddresses {
	return &MsgDeleteAddresses{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteAddresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
