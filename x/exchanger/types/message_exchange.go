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

var _ sdk.Msg = &MsgExchange{}

func NewMsgExchange(creator string, denom string, amount string, denomTo string) *MsgExchange {
	return &MsgExchange{
		Creator: creator,
		Denom:   denom,
		Amount:  amount,
		DenomTo: denomTo,
	}
}

func (msg *MsgExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// validate amount
	amount, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount (%s)", msg.Amount)
	}

	if !amount.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "amount must be positive")
	}

	// validate denoms
	if len(msg.Denom) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "denom cannot be empty")
	}
	if len(msg.DenomTo) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "denomTo cannot be empty")
	}

	return nil
}
