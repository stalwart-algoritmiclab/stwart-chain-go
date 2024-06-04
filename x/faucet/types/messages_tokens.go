package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTokens{}

func NewMsgCreateTokens(creator string, denom string, amount string) *MsgCreateTokens {
	return &MsgCreateTokens{
		Creator: creator,
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgCreateTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// check if Amount is parsable to math.Int
	_, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateTokens{}

func NewMsgUpdateTokens(creator string, id uint64, denom string, amount string) *MsgUpdateTokens {
	return &MsgUpdateTokens{
		Id:      id,
		Creator: creator,
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgUpdateTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// check if Amount is parsable to math.Int
	_, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteTokens{}

func NewMsgDeleteTokens(creator string, id uint64) *MsgDeleteTokens {
	return &MsgDeleteTokens{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
