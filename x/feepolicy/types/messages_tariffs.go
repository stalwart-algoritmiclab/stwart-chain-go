package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTariffs{}

func NewMsgCreateTariffs(
	creator string,
	denom string,
	tariff *Tariff,

) *MsgCreateTariffs {
	return &MsgCreateTariffs{
		Creator: creator,
		Denom:   denom,
		Tariffs: tariff,
	}
}

func (msg *MsgCreateTariffs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTariffs{}

func NewMsgUpdateTariffs(
	creator string,
	denom string,
	tariff *Tariff,

) *MsgUpdateTariffs {
	return &MsgUpdateTariffs{
		Creator: creator,
		Denom:   denom,
		Tariffs: tariff,
	}
}

func (msg *MsgUpdateTariffs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTariffs{}

func NewMsgDeleteTariffs(
	creator string,
	denom string,

) *MsgDeleteTariffs {
	return &MsgDeleteTariffs{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgDeleteTariffs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
