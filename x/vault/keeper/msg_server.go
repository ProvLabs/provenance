package keeper

import (
	"context"

	"github.com/provenance-io/provenance/x/vault/types"
)

// MsgServer is an alias for a Keeper that implements the vault.MsgServer interface.
type MsgServer struct {
	Keeper
}

// NewMsgServer returns a new instance of MsgServer.
func NewMsgServer(k Keeper) types.MsgServer {
	return MsgServer{
		Keeper: k,
	}
}

var _ types.MsgServer = MsgServer{}

// CreateVault creates a new vault.
func (k MsgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVaultRequest) (*types.MsgCreateVaultResponse, error) {
	panic("not implemented")
}

// Deposit deposits assets into a vault.
func (k MsgServer) Deposit(goCtx context.Context, msg *types.MsgDepositRequest) (*types.MsgDepositResponse, error) {
	panic("not implemented")
}

// Withdraw withdraws assets from a vault.
func (k MsgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdrawRequest) (*types.MsgWithdrawResponse, error) {
	panic("not implemented")
}

// Redeem redeems shares for underlying assets.
func (k MsgServer) Redeem(goCtx context.Context, msg *types.MsgRedeemRequest) (*types.MsgRedeemResponse, error) {
	panic("not implemented")
}
