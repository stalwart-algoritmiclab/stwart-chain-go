/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

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
