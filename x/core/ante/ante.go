/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package ante

import (
	sdkioerrors "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper, EVM Keeper and Fee Market Keeper.
type HandlerOptions struct {
	AccountKeeper          authante.AccountKeeper
	BankKeeper             types.BankKeeper
	FeeKeeper              BankKeeper
	ExtensionOptionChecker authante.ExtensionOptionChecker
	FeegrantKeeper         authante.FeegrantKeeper
	SignModeHandler        *txsigning.HandlerMap
	SigGasConsumer         authante.SignatureVerificationGasConsumer
	TxFeeChecker           authante.TxFeeChecker
	AccKeeper              AccountKeeper
	FeePolicyKeeper        FeePolicyKeeper
	CoreKeeper             CoreKeeper
	RefKeeper              RefKeeper
	RewardsKeeper          RewardsKeeper
	UsersKeeper            UserKeeper
	StakeKeeper            StakeKeeper
	StatsKeeper            StatsKeeper
}

func (options HandlerOptions) Validate() error {
	if options.AccountKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.SigGasConsumer == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "SigGasConsumer is required for AnteHandler")
	}
	if options.FeeKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "FeeKeeper is required for AnteHandler")
	}
	if options.FeePolicyKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "FeePolicyKeeper is required for AnteHandler")
	}
	if options.CoreKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "CoreKeeper is required for AnteHandler")
	}
	if options.RefKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "RefKeeper is required for AnteHandler")
	}
	if options.UsersKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "UsersKeeper is required for AnteHandler")
	}
	if options.RewardsKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "RewardsKeeper is required for AnteHandler")
	}
	if options.StakeKeeper == nil {
		return sdkioerrors.Wrap(sdkerrors.ErrLogic, "StakeKeeper is required for AnteHandler")
	}

	return nil
}

// NewAnteHandler returns an 'AnteHandler' that will run actions before a tx is sent to a module's handler.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if err := options.Validate(); err != nil {
		return nil, err
	}

	anteDecorators := []sdk.AnteDecorator{
		authante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		authante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),
		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		authante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		authante.NewIncrementSequenceDecorator(options.AccountKeeper),
		NewCoreDecorator(
			options.AccKeeper,
			options.FeeKeeper,
			options.FeePolicyKeeper,
			options.CoreKeeper,
			options.RefKeeper,
			options.UsersKeeper,
			options.RewardsKeeper,
			options.StakeKeeper,
			options.StatsKeeper,
		),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
