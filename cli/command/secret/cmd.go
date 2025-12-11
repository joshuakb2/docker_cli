package secret

import (
	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/joshuakb2/docker_cli/cli/command/completion"
	"github.com/joshuakb2/docker_cli/internal/commands"
	"github.com/joshuakb2/moby/client"
	"github.com/spf13/cobra"
)

func init() {
	commands.Register(newSecretCommand)
}

// newSecretCommand returns a cobra command for `secret` subcommands.
func newSecretCommand(dockerCLI command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Manage Swarm secrets",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCLI.Err()),
		Annotations: map[string]string{
			"version": "1.25",
			"swarm":   "manager",
		},
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newSecretListCommand(dockerCLI),
		newSecretCreateCommand(dockerCLI),
		newSecretInspectCommand(dockerCLI),
		newSecretRemoveCommand(dockerCLI),
	)
	return cmd
}

// completeNames offers completion for swarm secrets
func completeNames(dockerCLI completion.APIClientProvider) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		res, err := dockerCLI.Client().SecretList(cmd.Context(), client.SecretListOptions{})
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		var names []string
		for _, secret := range res.Items {
			names = append(names, secret.Spec.Name)
		}
		return names, cobra.ShellCompDirectiveNoFileComp
	}
}
