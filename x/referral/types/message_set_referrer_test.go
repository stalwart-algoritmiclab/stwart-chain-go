package types

import (
	"testing"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetReferrer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetReferrer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetReferrer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetReferrer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
