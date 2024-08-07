/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetReferrer{}

func NewMsgSetReferrer(creator string, referrerAddress string, referralAddress string) *MsgSetReferrer {
	return &MsgSetReferrer{
		Creator:         creator,
		ReferrerAddress: referrerAddress,
		ReferralAddress: referralAddress,
	}
}

func (msg *MsgSetReferrer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ReferrerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid referrer address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ReferralAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid referral address (%s)", err)
	}

	return nil
}
