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

var _ sdk.Msg = &MsgSend{}

func NewMsgSend(creator string, from string, to string, amount string, denom string) *MsgSend {
	return &MsgSend{
		Creator: creator,
		From:    from,
		To:      to,
		Amount:  amount,
		Denom:   denom,
	}
}

func (msg *MsgSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	addressTo, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address to: %s is not correct", addressTo)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if amount.IsNil() || !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", msg.Amount)
	}

	if amount.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount is zero")
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom is empty")
	}

	return nil
}
