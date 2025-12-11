package network

import (
	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/joshuakb2/docker_cli/internal/commands"
	"github.com/spf13/cobra"
)

func init() {
	commands.Register(newNetworkCommand)
}

// newNetworkCommand returns a cobra command for `network` subcommands
func newNetworkCommand(dockerCLI command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "network",
		Short:       "Manage networks",
		Args:        cli.NoArgs,
		RunE:        command.ShowHelp(dockerCLI.Err()),
		Annotations: map[string]string{"version": "1.21"},

		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newConnectCommand(dockerCLI),
		newCreateCommand(dockerCLI),
		newDisconnectCommand(dockerCLI),
		newInspectCommand(dockerCLI),
		newListCommand(dockerCLI),
		newRemoveCommand(dockerCLI),
		newPruneCommand(dockerCLI),
	)
	return cmd
}
