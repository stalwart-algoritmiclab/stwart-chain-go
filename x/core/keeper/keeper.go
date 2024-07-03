/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	exchangermoduletypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
	feepolicymoduletypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
	securedmodulekeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/keeper"
	securedmoduletypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
	systemrewardsmoduletypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/systemrewards/types"
	usersmodulekeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/keeper"
)

var _ CoreKeeper = (*Keeper)(nil)

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
	authority string,
	securedKeeper securedmodulekeeper.SecuredKeeper,
	userKeeper usersmodulekeeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	// init modules names and addresses list
	modulesAddresses := []string{
		types.ModuleName,
		exchangermoduletypes.ModuleName,
		securedmoduletypes.ModuleName,
		feepolicymoduletypes.ModuleName,
		systemrewardsmoduletypes.ModuleName,
	}

	modulesList := make([]types.ModuleInfo, 0, len(modulesAddresses))
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
