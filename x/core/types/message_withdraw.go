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

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
)

var _ sdk.Msg = &MsgWithdraw{}

func NewMsgWithdraw(creator string, amount string, denom string, address string) *MsgWithdraw {
	return &MsgWithdraw{
		Creator: creator,
		Amount:  amount,
		Denom:   denom,
		Address: address,
	}
}

func (msg *MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address: %s is not correct", address)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", msg.Amount)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom is empty")
	}

	coinAmount := sdk.NewCoin(msg.Denom, amount)

	if coinAmount.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", msg.Amount)
	}

	if coinAmount.Denom == domain.DenomStake {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cannot withdraw stake")
	}

	return nil
}
