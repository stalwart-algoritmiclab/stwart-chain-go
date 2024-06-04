package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	addressTo, _ := sdk.AccAddressFromBech32(msg.To)
	amount, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}

	coinAmount := sdk.NewCoin(msg.Denom, amount)
	coins := sdk.NewCoins(coinAmount)

	var moduleName string
	// check if from address exists
	for _, info := range k.modulesList {
		if msg.From == info.Address {
			moduleName = info.Name
			break
		}
	}

	if moduleName == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid from address (%s)", msg.From)
	}

	if err := k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, addressTo, coins); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgSendResponse{}, nil
}
