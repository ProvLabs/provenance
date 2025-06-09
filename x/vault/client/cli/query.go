package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/provenance-io/provenance/x/vault/types"
)

// CmdQuery returns the parent command for all vault module queries.
func CmdQuery() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Aliases:                    []string{"vault"},
		Short:                      "Querying commands for the vault module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	// Subcommands will be added here later.
	return cmd
}
