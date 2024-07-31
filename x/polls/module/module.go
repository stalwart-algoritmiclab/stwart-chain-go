package polls

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	// this line is used by starport scaffolding # 1

	modulev1 "github.com/stalwart-algoritmiclab/stwart-chain-go/api/stwartchain/polls/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
	securedmodulekeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/keeper"
)

const (
	FailureReasonNotEnoughVoters      = "not enough voters"
	FailureReasonNotEnoughVotesAmount = "not enough votes amount"
	FailedReasonNoWinner              = "no winner in this poll"
)

var (
	_ module.AppModuleBasic      = (*AppModule)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasInvariants       = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)

	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage.
// The default GenesisState need to be defined by the module developer and is primarily used for testing.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module.
// It should be incremented on each consensus-breaking change introduced by the module.
// To avoid wrong/empty versions, the initial version should be set to 1.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block.
// The begin block implementation is optional.
func (am AppModule) BeginBlock(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
	timeNow := ctx.BlockTime().UTC()

	for _, poll := range am.keeper.GetAllPollsByStatuses(ctx, keeper.StatusVotingPeriod, keeper.StatusPending) {
		votingEndTime, err := time.Parse(time.RFC3339, poll.VotingEndTime)
		if err != nil {
			err = errorsmod.Wrapf(sdkerrors.ErrLogic, "invalid voting end time in poll %d", poll.Id)
			logger.Error("beginBlock", "error", err)

			poll.Status = keeper.StatusFailed.String()
			poll.FailureReason = err.Error()
			am.keeper.SetPolls(ctx, poll)

			continue
		}

		votingStartTime, err := time.Parse(time.RFC3339, poll.VotingStartTime)
		if err != nil {
			err = errorsmod.Wrapf(sdkerrors.ErrLogic, "invalid voting start time in poll %d", poll.Id)
			logger.Error("beginBlock", "error", err)

			poll.Status = keeper.StatusFailed.String()
			poll.FailureReason = err.Error()
			am.keeper.SetPolls(ctx, poll)

			continue
		}

		switch poll.Status {
		case keeper.StatusVotingPeriod.String():
			if votingEndTime.Before(timeNow) {
				poll, err := am.countVotes(ctx, poll)
				if err != nil {
					logger.Error("beginBlock", "error", err)

					poll.Status = keeper.StatusFailed.String()
					poll.FailureReason = err.Error()
					am.keeper.SetPolls(ctx, poll)

					continue
				}

				am.keeper.SetPolls(ctx, poll)
			}
		case keeper.StatusPending.String():
			if votingStartTime.Before(timeNow) {
				poll.Status = keeper.StatusVotingPeriod.String()
				am.keeper.SetPolls(ctx, poll)
			}
		default:
			continue
		}
	}

	return nil
}

func (am AppModule) countVotes(ctx sdk.Context, poll types.Polls) (types.Polls, error) {
	defer func() {
		if poll.Status != keeper.StatusRejected.String() {
			if err := am.returnProposerDeposit(ctx, poll); err != nil {
				ctx.Logger().Error("countVotes", "return proposer deposit error", err)
			}
		}
	}()

	var (
		totalVotedAmount   sdk.Coins
		totalAccountsVoted uint64
		winningAmount      sdk.Coins
		winnerID           uint64 = 1
		maxVotesCount      int64  = 0
	)

	for _, option := range poll.Options {
		// count total votes stats
		totalVotedAmount = totalVotedAmount.Add(option.TokensAmount...)
		totalAccountsVoted += option.VotersCount

		// count winner stats
		if sdk.NewCoins(option.TokensAmount...).IsAllGT(winningAmount) {
			winningAmount = option.TokensAmount
			winnerID = option.Id
			maxVotesCount = 1

			continue
		}

		if sdk.NewCoins(option.TokensAmount...).Equal(winningAmount) {
			maxVotesCount++
		}
	}

	// validate poll minimum thresholds
	if totalAccountsVoted < poll.MinAddressesCount {
		poll.Status = keeper.StatusFailed.String()
		poll.FailureReason = FailureReasonNotEnoughVoters

		return poll, nil
	}

	if sdk.NewCoins(totalVotedAmount...).IsAllLT(poll.MinVotedCoinsAmount) {
		poll.Status = keeper.StatusFailed.String()
		poll.FailureReason = FailureReasonNotEnoughVotesAmount

		return poll, nil
	}

	// found winner in poll, if not found - fail poll
	winner, found := am.keeper.GetPollsOptionByID(poll, winnerID)
	if !found {
		poll.Status = keeper.StatusFailed.String()
		poll.FailureReason = FailedReasonNoWinner

		return poll, nil
	}

	// if winner is veto, burn proposer deposit and set poll status to rejected
	if winner.IsVeto {
		if err := am.burnProposerDeposit(ctx); err != nil {
			return poll, err
		}

		poll.Status = keeper.StatusRejected.String()
		return poll, nil
	}

	if maxVotesCount > 1 {
		poll.Status = keeper.StatusFailed.String()
		poll.FailureReason = FailedReasonNoWinner

		return poll, nil
	}

	poll.Status = keeper.StatusPassed.String()
	winner.IsWinner = true
	keeper.UpdateOption(poll, winner)

	return poll, nil
}

func (am AppModule) burnProposerDeposit(ctx sdk.Context) error {
	moduleParams, ok := am.keeper.GetPollsParams(ctx)
	if !ok {
		return errorsmod.Wrap(sdkerrors.ErrAppConfig, "polls module params are not set")
	}
	if !moduleParams.BurnVeto {
		return nil
	}

	if err := am.bankKeeper.BurnCoins(ctx, types.ModuleName, moduleParams.ProposerDeposit); err != nil {
		return err
	}

	return nil
}

func (am AppModule) returnProposerDeposit(ctx sdk.Context, poll types.Polls) error {
	moduleParams, ok := am.keeper.GetPollsParams(ctx)
	if !ok {
		return errorsmod.Wrap(sdkerrors.ErrAppConfig, "polls module params are not set")
	}

	proposerAddress, err := sdk.AccAddressFromBech32(poll.ProposerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrLogic, "invalid proposer address in poll %d", poll.Id)
	}

	if err := am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, proposerAddress, moduleParams.ProposerDeposit); err != nil {
		return err
	}

	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block.
// The end block implementation is optional.
func (am AppModule) EndBlock(_ context.Context) error {
	return nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *modulev1.Module
	Logger       log.Logger

	SecuredKeeper securedmodulekeeper.SecuredKeeper
	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	StakingKeeper types.StakingKeeper
}

type ModuleOutputs struct {
	depinject.Out

	PollsKeeper keeper.Keeper
	Module      appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Logger,
		authority.String(),
		in.SecuredKeeper,
		in.StakingKeeper,
		in.BankKeeper,
		in.AccountKeeper,
	)
	m := NewAppModule(
		in.Cdc,
		k,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return ModuleOutputs{PollsKeeper: k, Module: m}
}
