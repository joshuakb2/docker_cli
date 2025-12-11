package trust

import (
	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/spf13/cobra"
)

// newTrustKeyCommand returns a cobra command for `trust key` subcommands
func newTrustKeyCommand(dockerCLI command.Streams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "Manage keys for signing Docker images",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCLI.Err()),

		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newKeyGenerateCommand(dockerCLI),
		newKeyLoadCommand(dockerCLI),
	)
	return cmd
}
