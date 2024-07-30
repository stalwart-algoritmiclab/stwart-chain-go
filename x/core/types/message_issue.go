/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssue{}

func NewMsgIssue(creator string, amount string, denom string, address string) *MsgIssue {
	return &MsgIssue{
		Creator: creator,
		Amount:  amount,
		Denom:   denom,
		Address: address,
	}
}

func (msg *MsgIssue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address: %s is not correct", msg.Address)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if amount.IsNil() || !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom is empty")
	}

	return nil
}
