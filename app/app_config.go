package app

import (
	"time"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	authzmodulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	circuitmodulev1 "cosmossdk.io/api/cosmos/circuit/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	govmodulev1 "cosmossdk.io/api/cosmos/gov/module/v1"
	groupmodulev1 "cosmossdk.io/api/cosmos/group/module/v1"
	mintmodulev1 "cosmossdk.io/api/cosmos/mint/module/v1"
	nftmodulev1 "cosmossdk.io/api/cosmos/nft/module/v1"
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	vestingmodulev1 "cosmossdk.io/api/cosmos/vesting/module/v1"
	"cosmossdk.io/core/appconfig"
	circuittypes "cosmossdk.io/x/circuit/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	"cosmossdk.io/x/nft"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"google.golang.org/protobuf/types/known/durationpb"

	coremodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/core/module"
	exchangermodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/exchanger/module"
	faucetmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/faucet/module"
	feepolicymodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/feepolicy/module"
	ratesmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/rates/module"
	referralmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/referral/module"
	securedmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/secured/module"
	stakemodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/stake/module"
	statsmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/stats/module"
	stwartmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/stwart/module"
	systemrewardsmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/systemrewards/module"
	usersmodulev1 "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/api/stwartchain/users/module"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/module" // import for side-effects
	coremoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/module" // import for side-effects
	exchangermoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/module" // import for side-effects
	faucetmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/module" // import for side-effects
	feepolicymoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/module" // import for side-effects
	ratesmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/module" // import for side-effects
	referralmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/module" // import for side-effects
	securedmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stake/module" // import for side-effects
	stakemoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stake/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/module" // import for side-effects
	statsmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stwart/module" // import for side-effects
	stwartmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stwart/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/module" // import for side-effects
	systemrewardsmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"
	_ "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/module" // import for side-effects
	usersmoduletypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"
	// this line is used by starport scaffolding # stargate/app/moduleImport
)

var (
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	genesisModuleOrder = []string{
		// cosmos-sdk/ibc modules
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		circuittypes.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		consensustypes.ModuleName,
		// chain modules
		stwartmoduletypes.ModuleName,
		securedmoduletypes.ModuleName,
		feepolicymoduletypes.ModuleName,
		systemrewardsmoduletypes.ModuleName,
		coremoduletypes.ModuleName,
		referralmoduletypes.ModuleName,
		usersmoduletypes.ModuleName,
		statsmoduletypes.ModuleName,
		stakemoduletypes.ModuleName,
		ratesmoduletypes.ModuleName,
		faucetmoduletypes.ModuleName,
		exchangermoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockers = []string{
		// cosmos sdk modules
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authz.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		capabilitytypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		stwartmoduletypes.ModuleName,
		securedmoduletypes.ModuleName,
		feepolicymoduletypes.ModuleName,
		systemrewardsmoduletypes.ModuleName,
		coremoduletypes.ModuleName,
		referralmoduletypes.ModuleName,
		usersmoduletypes.ModuleName,
		statsmoduletypes.ModuleName,
		ratesmoduletypes.ModuleName,
		stakemoduletypes.ModuleName,
		faucetmoduletypes.ModuleName,
		exchangermoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	}

	endBlockers = []string{
		// cosmos sdk modules
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		capabilitytypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		stwartmoduletypes.ModuleName,
		securedmoduletypes.ModuleName,
		feepolicymoduletypes.ModuleName,
		systemrewardsmoduletypes.ModuleName,
		coremoduletypes.ModuleName,
		referralmoduletypes.ModuleName,
		usersmoduletypes.ModuleName,
		statsmoduletypes.ModuleName,
		stakemoduletypes.ModuleName,
		ratesmoduletypes.ModuleName,
		faucetmoduletypes.ModuleName,
		exchangermoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
	}

	preBlockers = []string{
		upgradetypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/preBlockers
	}

	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: nft.ModuleName},
		{Account: ibctransfertypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: ibcfeetypes.ModuleName},
		{Account: icatypes.ModuleName},
		{Account: securedmoduletypes.ModuleName},
		{Account: feepolicymoduletypes.ModuleName},
		{Account: coremoduletypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: systemrewardsmoduletypes.ModuleName},
		{Account: statsmoduletypes.ModuleName},
		{Account: ratesmoduletypes.ModuleName},
		{Account: faucetmoduletypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: exchangermoduletypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},

		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		nft.ModuleName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// appConfig application configuration (used by depinject)
	appConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       Name,
					PreBlockers:   preBlockers,
					BeginBlockers: beginBlockers,
					EndBlockers:   endBlockers,
					InitGenesis:   genesisModuleOrder,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: genesisModuleOrder,
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: nil,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             AccountAddressPrefix,
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   nft.ModuleName,
				Config: appconfig.WrapAny(&nftmodulev1.Module{}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name: stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{
					// NOTE: specifying a prefix is only necessary when using bech32 addresses
					// If not specfied, the auth Bech32Prefix appended with "valoper" and "valcons" is used by default
					Bech32PrefixValidator: AccountAddressPrefix + "valoper",
					Bech32PrefixConsensus: AccountAddressPrefix + "valcons",
				}),
			},
			{
				Name:   slashingtypes.ModuleName,
				Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   minttypes.ModuleName,
				Config: appconfig.WrapAny(&mintmodulev1.Module{}),
			},
			{
				Name: group.ModuleName,
				Config: appconfig.WrapAny(&groupmodulev1.Module{
					MaxExecutionPeriod: durationpb.New(time.Second * 1209600),
					MaxMetadataLen:     255,
				}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   circuittypes.ModuleName,
				Config: appconfig.WrapAny(&circuitmodulev1.Module{}),
			},
			{
				Name:   stwartmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&stwartmodulev1.Module{}),
			},
			{
				Name:   securedmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&securedmodulev1.Module{}),
			},
			{
				Name:   feepolicymoduletypes.ModuleName,
				Config: appconfig.WrapAny(&feepolicymodulev1.Module{}),
			},
			{
				Name:   systemrewardsmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&systemrewardsmodulev1.Module{}),
			},
			{
				Name:   coremoduletypes.ModuleName,
				Config: appconfig.WrapAny(&coremodulev1.Module{}),
			},
			{
				Name:   referralmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&referralmodulev1.Module{}),
			},
			{
				Name:   usersmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&usersmodulev1.Module{}),
			},
			{
				Name:   statsmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&statsmodulev1.Module{}),
			},
			{
				Name:   ratesmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&ratesmodulev1.Module{}),
			},
			{
				Name:   stakemoduletypes.ModuleName,
				Config: appconfig.WrapAny(&stakemodulev1.Module{}),
			},
			{
				Name:   faucetmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&faucetmodulev1.Module{}),
			},
			{
				Name:   exchangermoduletypes.ModuleName,
				Config: appconfig.WrapAny(&exchangermodulev1.Module{}),
			},
			// this line is used by starport scaffolding # stargate/app/moduleConfig
		},
	})
)