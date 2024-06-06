/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	exchangermoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
	faucetmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
	securedmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
	stwartmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stwart/types"
)

var (
	maccPerms = map[string][]string{
		// Native SDK module accounts
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		icatypes.ModuleName:            nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},

		// Custom module accounts
		exchangermoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		faucetmoduletypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		securedmoduletypes.ModuleName:   {authtypes.Minter, authtypes.Burner},
		stwartmoduletypes.ModuleName:    nil,
	}

	blockedModuleAccounts = map[string]bool{
		authtypes.FeeCollectorName:     false,
		distrtypes.ModuleName:          false,
		stakingtypes.BondedPoolName:    false,
		stakingtypes.NotBondedPoolName: false,
		ibctransfertypes.ModuleName:    false,
		icatypes.ModuleName:            false,
	}
)

func moduleAccToAddress[V any](accs map[string]V) map[string]bool {
	addrs := make(map[string]bool)
	for acc := range accs {
		addrs[authtypes.NewModuleAddress(acc).String()] = true
	}
	return addrs
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	// By default, returns all the app's blocked module account addresses.
	// Other regular addresses can also be added here.
	return moduleAccToAddress(blockedModuleAccounts)
}
