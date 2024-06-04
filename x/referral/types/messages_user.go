package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgCreateUser(
	creator string,
	accountAddress string,
	referrer string,
	referrals []string,

) *MsgCreateUser {
	return &MsgCreateUser{
		Creator:        creator,
		AccountAddress: accountAddress,
		Referrer:       referrer,
		Referrals:      referrals,
	}
}

func (msg *MsgCreateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUser{}

func NewMsgUpdateUser(
	creator string,
	accountAddress string,
	referrer string,
	referrals []string,

) *MsgUpdateUser {
	return &MsgUpdateUser{
		Creator:        creator,
		AccountAddress: accountAddress,
		Referrer:       referrer,
		Referrals:      referrals,
	}
}

func (msg *MsgUpdateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUser{}

func NewMsgDeleteUser(
	creator string,
	accountAddress string,

) *MsgDeleteUser {
	return &MsgDeleteUser{
		Creator:        creator,
		AccountAddress: accountAddress,
	}
}

func (msg *MsgDeleteUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
