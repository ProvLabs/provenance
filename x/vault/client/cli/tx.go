package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/provenance-io/provenance/x/vault/types"
)

// TxCmd returns the parent command for all vault module transactions.
func TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{"vault"},
		Short:                      "Transaction commands for the vault module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Add subcommands here in the future.
	return cmd
}
