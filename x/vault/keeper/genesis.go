package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/provenance-io/provenance/x/vault/types"
)

// InitGenesis initializes the vault module's state from a genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	if genState == nil {
		return
	}

	// k.SetParams(ctx, genState.Params)

	// store := k.getStore(ctx)
	// for i, vault := range genState.Vaults {
	// 	if err := k.setVaultInStore(store, &vault); err != nil {
	// 		panic(fmt.Errorf("failed to store Vaults[%d]: %w", i, err))
	// 	}
	// }
}

// ExportGenesis exports the current state as a genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	// store := k.getStore(ctx)
	// genState := &types.GenesisState{
	// 	Params: k.GetParams(ctx),
	// }

	// k.IterateVaults(ctx, func(vault *types.Vault) bool {
	// 	genState.Vaults = append(genState.Vaults, *vault)
	// 	return false
	// })

	return &types.GenesisState{}
}
