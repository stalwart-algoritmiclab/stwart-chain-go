package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRefReward{}

func NewMsgRefReward(creator string, amount sdk.Coin, referrer string) *MsgRefReward {
	return &MsgRefReward{
		Creator:  creator,
		Amount:   amount,
		Referrer: referrer,
	}
}

func (msg *MsgRefReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
