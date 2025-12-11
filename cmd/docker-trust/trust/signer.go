package trust

import (
	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/spf13/cobra"
)

// newTrustSignerCommand returns a cobra command for `trust signer` subcommands
func newTrustSignerCommand(dockerCLI command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signer",
		Short: "Manage entities who can sign Docker images",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCLI.Err()),

		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newSignerAddCommand(dockerCLI),
		newSignerRemoveCommand(dockerCLI),
	)
	return cmd
}
