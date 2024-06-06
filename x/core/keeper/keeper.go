/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"
	exchangermoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
	feepolicymoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
	securedmodulekeeper "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/keeper"
	securedmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
	systemrewardsmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"
	usersmodulekeeper "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/keeper"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		modulesList   []types.ModuleInfo
		bankKeeper    types.BankKeeper
		userKeeper    usersmodulekeeper.Keeper
		securedKeeper securedmodulekeeper.SecuredKeeper
		accountKeeper types.AccountKeeper
		authority     string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	bankKeeper types.BankKeeper,
	securedKeeper securedmodulekeeper.SecuredKeeper,
	accountKeeper types.AccountKeeper,
	userKeeper usersmodulekeeper.Keeper,
	authority string,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	// init modules names and addresses list
	modulesAddresses := []string{
		types.ModuleName,
		feepolicymoduletypes.ModuleName,
		securedmoduletypes.ModuleName,
		systemrewardsmoduletypes.ModuleName,
		exchangermoduletypes.ModuleName,
	}

	var modulesList []types.ModuleInfo
	for _, s := range modulesAddresses {
		addr := accountKeeper.GetModuleAddress(s)
		if addr == nil {
			panic("address for the module not found: " + s)
		}
		info := types.ModuleInfo{
			Name:    s,
			Address: addr.String(),
		}

		modulesList = append(modulesList, info)
	}

	return Keeper{
		cdc:           cdc,
		storeService:  storeService,
		authority:     authority,
		bankKeeper:    bankKeeper,
		securedKeeper: securedKeeper,
		accountKeeper: accountKeeper,
		userKeeper:    userKeeper,
		modulesList:   modulesList,
		logger:        logger,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Burn burns coins from the void module account.
func (k Keeper) Burn(ctx sdk.Context, address sdk.AccAddress, amount sdk.Coin) error {
	if amount.IsZero() {
		return nil
	}

	if amount.IsNegative() {
		return fmt.Errorf("invalid amount: %s", amount)
	}

	if amount.Denom == domain.DenomStake {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "cannot burn stake")
	}

	coins := sdk.NewCoins(amount)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, coins); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return err
	}

	k.AddBurnedToDailyStats(ctx, coins...)

	msg := &types.MsgBurn{
		Creator: types.ModuleName,
		Amount:  amount.Amount.Uint64(),
		Denom:   amount.Denom,
		Address: address.String(),
	}
	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return err
	}

	return nil
}
