/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package ante

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	feepolicytypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
	referralmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"
)

// AccountKeeper specifies the interface that FeeDeductor requires
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
	GetParams(ctx context.Context) (params types.Params)
	SetAccount(ctx context.Context, acc sdk.AccountI)
}

// BankKeeper specifies the interface that FeeDeductor requires
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// CoreKeeper specifies the interface that FeeDeductor requires for core module
type CoreKeeper interface {
	Burn(ctx sdk.Context, address sdk.AccAddress, amount sdk.Coin) error
	// SetStatsNoFee(ctx context.Context, amount sdk.Coins)
	// SetStatsFee(ctx context.Context, amountFee sdk.Coins, amountTx sdk.Coins)
}

// FeePolicyKeeper specifies the interface that FeeDeductor requires
type FeePolicyKeeper interface {
	GetAddress(ctx sdk.Context, address string) (val feepolicytypes.Address, found bool)
	GetFees(ctx sdk.Context, receiverAssets sdk.Coins, sendAmount sdk.Coin) (val feepolicytypes.Fees, minRefAmount sdk.Coin, found bool)
}

// RewardsKeeper specifies the interface that FeeDeductor requires
type RewardsKeeper interface {
	AddStats(ctx context.Context, coins ...sdk.Coin) // Sys Reward Stats
}

// RefKeeper specifies the interface that FeeDeductor requires
type RefKeeper interface {
	GetUser(ctx context.Context, accountAddress string) (val referralmoduletypes.User, found bool)
}

type UserKeeper interface {
	CountUsers(ctx context.Context, userAddresses []string) uint64
	IncrementTotalUsers(ctx context.Context)
	AddNewUserToStat(ctx context.Context)
}

type StakeKeeper interface {
	GetFreeStake(ctx context.Context, address sdk.AccAddress) sdkmath.Int
}
