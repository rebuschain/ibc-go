package cli

import (
	"github.com/spf13/cobra"

	controllercli "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/controller/client/cli"
	hostcli "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/host/client/cli"
)

// GetQueryCmd returns the query commands for the interchain-accounts submodule
func GetQueryCmd() *cobra.Command {
	icaQueryCmd := &cobra.Command{
		Use:                        "interchain-accounts",
		Aliases:                    []string{"ica"},
		Short:                      "interchain-accounts subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	icaQueryCmd.AddCommand(
		controllercli.GetQueryCmd(),
		hostcli.GetQueryCmd(),
	)

	return icaQueryCmd
}

// NewTxCmd returns the tx commands for the interchain-accounts submodule
func NewTxCmd() *cobra.Command {
	icaTxCmd := &cobra.Command{
		Use:                        "interchain-accounts",
		Aliases:                    []string{"ica"},
		Short:                      "interchain-accounts tx subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	icaTxCmd.AddCommand(
		controllercli.NewTxCmd(),
	)

	return icaTxCmd
}
