package types

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the vault module.
	ModuleName = "vault"

	// StoreKey is the store key string for the vault module.
	StoreKey = ModuleName
)

// GetVaultAddress returns the module account address for the given vaultID.
func GetVaultAddress(vaultID uint32) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(fmt.Sprintf("%s/%d", ModuleName, vaultID))))
}
