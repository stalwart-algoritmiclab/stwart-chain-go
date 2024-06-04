package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgFees{}

func NewMsgFees(comission sdk.Coin, addressTo string) *MsgFees {
	return &MsgFees{
		Comission: comission,
		AddressTo: addressTo,
	}
}

func (msg *MsgFees) ValidateBasic() error {
	return nil
}
