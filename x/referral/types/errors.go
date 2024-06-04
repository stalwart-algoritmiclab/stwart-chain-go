package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/referral module sentinel errors
var (
	ErrInvalidSigner      = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample             = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidDate        = sdkerrors.Register(ModuleName, 1102, "invalid date")
	ErrInvalidAddress     = sdkerrors.Register(ModuleName, 1103, "invalid address")
	ErrReferrerAlreadySet = sdkerrors.Register(ModuleName, 1104, "referrer already set")
	ErrAccountNotFound    = sdkerrors.Register(ModuleName, 1105, "account not found")
)
