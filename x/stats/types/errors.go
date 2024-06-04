package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

// x/stats module sentinel errors
var (
	ErrInvalidDate   = sdkerrors.Register(ModuleName, 2, "invalid date")
	ErrNotFound      = sdkerrors.Register(ModuleName, uint32(codes.NotFound), "not found %s")
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
)
