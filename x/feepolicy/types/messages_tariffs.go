/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import (
	"strconv"

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

	if err = validateTariff(msg.Tariffs); err != nil {
		return err
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

	if err = validateTariff(msg.Tariffs); err != nil {
		return err
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

func validateTariff(tariff *Tariff) error {
	tariffAmount, err := strconv.ParseFloat(tariff.Amount, 64)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", tariff.Amount)
	}

	if tariffAmount < 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff amount must be greater than zero, current value (%s)", tariff.Amount)

	}

	tariffMinRefBalance, err := strconv.ParseFloat(tariff.MinRefBalance, 64)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid min ref balance (%s)", tariff.MinRefBalance)
	}

	if tariffMinRefBalance < 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff min ref balance must be greater than zero, current value (%s)", tariff.MinRefBalance)

	}

	for _, fee := range tariff.Fees {
		tariffFee, err := strconv.ParseFloat(fee.Fee, 64)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fee (%s), fee id: (%d)", fee.Fee, fee.Id)
		}

		if tariffFee < 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff fee must be greater than zero, current value (%s), fee id: (%d)", fee.Fee, fee.Id)

		}

		tariffRefReward, err := strconv.ParseFloat(fee.RefReward, 64)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ref reward (%s), fee id: (%d)", fee.RefReward, fee.Id)
		}

		if tariffRefReward < 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff ref reward must be greater than zero, current value (%s), fee id: (%d)", fee.RefReward, fee.Id)
		}

		tariffStakeReward, err := strconv.ParseFloat(fee.StakeReward, 64)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid stake reward (%s), fee id: (%d)", fee.StakeReward, fee.Id)
		}

		if tariffStakeReward < 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff stake reward must be greater than zero, current value (%s), fee id: (%d)", fee.StakeReward, fee.Id)
		}

		if fee.MinAmount < 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff min amount must be greater than zero, current value (%s), fee id: (%d)", fee.MinAmount, fee.Id)
		}

		tariffAmountFrom, err := strconv.ParseFloat(fee.AmountFrom, 64)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount from (%s), fee id: (%d)", fee.RefReward, fee.Id)
		}

		if tariffAmountFrom < 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tariff amount from must be greater than zero, current value (%s), fee id: (%d)", fee.RefReward, fee.Id)
		}
	}

	return nil
}
