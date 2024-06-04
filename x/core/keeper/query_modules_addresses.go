package keeper

import (
	"context"
	"fmt"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

// ModulesAddresses returns all modules addresses
func (k Keeper) ModulesAddresses(_ context.Context, _ *types.QueryModulesAddressesRequest) (*types.QueryModulesAddressesResponse, error) {
	resp := []string{}
	for _, s := range k.modulesList {
		resp = append(resp, fmt.Sprintf("%s: %s", s.Name, s.Address))
	}

	return &types.QueryModulesAddressesResponse{Address: resp}, nil
}
