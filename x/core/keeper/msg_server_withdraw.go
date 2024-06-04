package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	address, _ := sdk.AccAddressFromBech32(msg.Address)
	acc := k.Keeper.accountKeeper.GetAccount(ctx, address)
	if acc == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "address %s is not found", msg.Address)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount (%s)", msg.Amount)
	}

	coinAmount := sdk.NewCoin(msg.Denom, amount)
	coins := sdk.NewCoins(coinAmount)

	if coinAmount.Denom == domain.DenomStake {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "cannot withdraw stake")
	}

	if err := k.Burn(ctx, address, coinAmount); err != nil {
		return nil, err
	}

	k.AddWithdrawnToDailyStats(ctx, coins...)
	k.AddBurnedToDailyStats(ctx, coins...)

	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
