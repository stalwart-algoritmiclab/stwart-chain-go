package keeper

import (
	"context"

	sdkioerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"
)

// SetUser set a specific user in the store from its index
func (k Keeper) SetUser(ctx context.Context, user types.User) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKeyPrefix))
	b := k.cdc.MustMarshal(&user)
	store.Set(types.UserKey(
		user.AccountAddress,
	), b)
}

// GetUser returns a user from its index
func (k Keeper) GetUser(ctx context.Context, accountAddress string) (val types.User, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKeyPrefix))

	b := store.Get(types.UserKey(
		accountAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUser removes a user from the store
func (k Keeper) RemoveUser(ctx context.Context, accountAddress string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKeyPrefix))
	store.Delete(types.UserKey(
		accountAddress,
	))
}

// GetAllUser returns all user
func (k Keeper) GetAllUser(ctx context.Context) (list []types.User) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.User
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetReferrer set a specific referrer in user state.
func (k Keeper) SetReferrer(goCtx context.Context, msg *types.MsgSetReferrer) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := sdk.AccAddressFromBech32(msg.ReferralAddress); err != nil {
		return sdkioerrors.Wrapf(types.ErrInvalidAddress, "referral account address: %s", err.Error())
	}

	if _, err := sdk.AccAddressFromBech32(msg.ReferrerAddress); err != nil {
		return sdkioerrors.Wrapf(types.ErrInvalidAddress, "referrer account address: %s", err.Error())
	}

	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkioerrors.Wrapf(types.ErrInvalidAddress, "creator account address: %s", err.Error())
	}

	if msg.ReferralAddress != msg.Creator {
		_, found := k.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
		if !found {
			return sdkioerrors.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
		}
	}

	if msg.ReferrerAddress == msg.ReferralAddress {
		return sdkioerrors.Wrapf(sdkerrors.ErrInvalidRequest, "referrer and referral can't be the same")
	}

	// Referrer and referral may not exist in account keeper

	referralUser, found := k.GetUser(ctx, msg.ReferralAddress)
	if !found {
		referralUser = types.User{
			AccountAddress: msg.ReferralAddress,
			Referrer:       msg.ReferrerAddress,
			Referrals:      nil,
		}
	} else {
		if referralUser.Referrer != "" {
			return sdkioerrors.Wrapf(types.ErrReferrerAlreadySet, "referral: %s, referrer: %s",
				msg.ReferralAddress, msg.ReferrerAddress)
		}

		if isExistsInSlice(referralUser.Referrals, msg.ReferrerAddress) {
			return sdkioerrors.Wrapf(sdkerrors.ErrConflict, "referrer %s is already exist in user's referrals", msg.ReferrerAddress)
		}

		referralUser.Referrer = msg.ReferrerAddress
	}

	referrerUser, found := k.GetUser(ctx, msg.ReferrerAddress)
	if !found {
		referrerUser = types.User{
			AccountAddress: msg.ReferrerAddress,
			Referrer:       "",
			Referrals:      []string{msg.ReferralAddress},
		}
	} else {
		referrerUser.Referrals = append(referrerUser.Referrals, msg.ReferralAddress)
	}

	k.SetUser(ctx, referralUser)
	k.SetUser(ctx, referrerUser)

	return nil
}
