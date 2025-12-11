package volume

import (
	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/joshuakb2/docker_cli/internal/commands"
	"github.com/spf13/cobra"
)

func init() {
	commands.Register(newVolumeCommand)
}

// newVolumeCommand returns a cobra command for `volume` subcommands
func newVolumeCommand(dockerCLI command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "volume COMMAND",
		Short:       "Manage volumes",
		Args:        cli.NoArgs,
		RunE:        command.ShowHelp(dockerCLI.Err()),
		Annotations: map[string]string{"version": "1.21"},

		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newCreateCommand(dockerCLI),
		newInspectCommand(dockerCLI),
		newListCommand(dockerCLI),
		newRemoveCommand(dockerCLI),
		newPruneCommand(dockerCLI),
		newUpdateCommand(dockerCLI),
	)
	return cmd
}
