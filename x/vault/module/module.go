package module

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	// "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/core/appmodule"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/provenance-io/provenance/x/vault/keeper"

	// "github.com/provenance-io/provenance/x/vault/simulation"
	"github.com/provenance-io/provenance/x/vault/client/cli"
	"github.com/provenance-io/provenance/x/vault/types"
)

var (
	_ module.AppModuleBasic = (*AppModule)(nil)
	// _ module.AppModuleSimulation = (*AppModule)(nil)
	_ appmodule.AppModule = (*AppModule)(nil)
)

// AppModuleBasic defines the basic application module used by the vault module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the vault module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// DefaultGenesis returns default genesis state as raw bytes for the vault module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the vault module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return data.Validate()
}

// GetQueryCmd returns the CLI query commands for the vault module.
func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.CmdQuery()
}

// GetTxCmd returns the CLI transaction commands for the vault module.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.TxCmd()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the vault module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// RegisterInterfaces registers the vault module's interface types.
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// RegisterLegacyAminoCodec registers the vault module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// AppModule implements an application module for the vault module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(cdc codec.Codec, vaultKeeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         vaultKeeper,
	}
}

// IsOnePerModuleType indicates that this is the only vault module of its type in the app.
func (AppModule) IsOnePerModuleType() {}

// IsAppModule marks this as a Cosmos SDK AppModule.
func (AppModule) IsAppModule() {}

// RegisterInvariants registers the vault module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs genesis initialization for the vault module.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	// var genesisState vault.GenesisState
	// cdc.MustUnmarshalJSON(data, &genesisState)
	// am.keeper.InitGenesis(ctx, &genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// RegisterServices registers module services for gRPC Msg and Query handling.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(am.keeper))
}

// ConsensusVersion returns the consensus version of the vault module.
func (AppModule) ConsensusVersion() uint64 {
	return 1
}

// GenerateGenesisState generates a randomized genesis state for simulation.
func (am AppModule) GenerateGenesisState(simState *module.SimulationState) {
	// simulation.RandomizedGenState(simState)
}

// RandomizedParams returns randomized parameter changes for simulation.
func (AppModule) RandomizedParams(_ *rand.Rand) []simtypes.LegacyParamChange {
	return nil
}

// RegisterStoreDecoder registers the store decoder for simulation.
func (am AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	// sdr[types.StoreKey] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the simulation operations for the vault module.
func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	// return simulation.WeightedOperations()
	return []simtypes.WeightedOperation{}
}
