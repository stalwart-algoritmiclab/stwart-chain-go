package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/core module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidDate   = sdkerrors.Register(ModuleName, 2, "invalid date")
)