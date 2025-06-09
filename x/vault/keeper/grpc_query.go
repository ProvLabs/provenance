package keeper

import (
	"context"

	"github.com/provenance-io/provenance/x/vault/types"
)

// QueryServer is an alias for a Keeper that implements the vault.QueryServer interface.
type QueryServer struct {
	Keeper
}

// NewQueryServer returns a new instance of QueryServer.
func NewQueryServer(k Keeper) types.QueryServer {
	return QueryServer{Keeper: k}
}

var _ types.QueryServer = QueryServer{}

// Vaults returns a paginated list of all vaults.
func (k QueryServer) Vaults(goCtx context.Context, req *types.QueryVaultsRequest) (*types.QueryVaultsResponse, error) {
	panic("not implemented")
}

// Vault returns the configuration and state of a specific vault.
func (k QueryServer) Vault(goCtx context.Context, req *types.QueryVaultRequest) (*types.QueryVaultResponse, error) {
	panic("not implemented")
}

// TotalAssets returns the total amount of the underlying asset managed by the vault.
func (k QueryServer) TotalAssets(goCtx context.Context, req *types.QueryTotalAssetsRequest) (*types.QueryTotalAssetsResponse, error) {
	panic("not implemented")
}
