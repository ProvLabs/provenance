package types

import (
	cerrs "cosmossdk.io/errors"
)

var (
	ErrVaultNotFound = cerrs.Register(ModuleName, 2, "vault not found")
)
