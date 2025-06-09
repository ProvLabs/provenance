package types

// DefaultGenesisState returns the default genesis state for the vault module.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate checks that the Vault struct is well-formed.
func (v GenesisState) Validate() error {
	return nil
}
