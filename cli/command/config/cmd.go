package config

import (
	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/joshuakb2/docker_cli/cli/command/completion"
	"github.com/joshuakb2/docker_cli/internal/commands"
	"github.com/joshuakb2/moby/client"
	"github.com/spf13/cobra"
)

func init() {
	commands.Register(newConfigCommand)
}

// newConfigCommand returns a cobra command for `config` subcommands
func newConfigCommand(dockerCLI command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage Swarm configs",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCLI.Err()),
		Annotations: map[string]string{
			"version": "1.30",
			"swarm":   "manager",
		},
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newConfigListCommand(dockerCLI),
		newConfigCreateCommand(dockerCLI),
		newConfigInspectCommand(dockerCLI),
		newConfigRemoveCommand(dockerCLI),
	)
	return cmd
}

// completeNames offers completion for swarm configs
func completeNames(dockerCLI completion.APIClientProvider) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		res, err := dockerCLI.Client().ConfigList(cmd.Context(), client.ConfigListOptions{})
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		var names []string
		for _, config := range res.Items {
			names = append(names, config.ID)
		}
		return names, cobra.ShellCompDirectiveNoFileComp
	}
}
