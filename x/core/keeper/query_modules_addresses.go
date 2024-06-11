/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"fmt"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

// ModulesAddresses returns all modules addresses
func (k Keeper) ModulesAddresses(
	_ context.Context,
	_ *types.QueryModulesAddressesRequest,
) (*types.QueryModulesAddressesResponse, error) {
	resp := make([]string, 0, len(k.modulesList))
	for _, s := range k.modulesList {
		resp = append(resp, fmt.Sprintf("%s: %s", s.Name, s.Address))
	}

	return &types.QueryModulesAddressesResponse{Address: resp}, nil
}

// ModuleAddressesByName returns module address by name
func (k Keeper) ModuleAddressesByName(name string) (string, error) {
	for _, s := range k.modulesList {
		if s.Name == name {
			return s.Address, nil
		}
	}

	return "", fmt.Errorf("module %s not found", name)
}
