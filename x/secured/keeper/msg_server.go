/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
)

var _ SecuredKeeper = (*msgServer)(nil)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the msgServer interface for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
