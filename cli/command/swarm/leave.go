package swarm

import (
	"fmt"

	"github.com/joshuakb2/docker_cli/cli"
	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/joshuakb2/moby/client"
	"github.com/spf13/cobra"
)

func newLeaveCommand(dockerCLI command.Cli) *cobra.Command {
	var opts client.SwarmLeaveOptions

	cmd := &cobra.Command{
		Use:   "leave [OPTIONS]",
		Short: "Leave the swarm",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := dockerCLI.Client().SwarmLeave(cmd.Context(), opts); err != nil {
				return err
			}

			_, _ = fmt.Fprintln(dockerCLI.Out(), "Node left the swarm.")
			return nil
		},
		Annotations: map[string]string{
			"version": "1.24",
			"swarm":   "active",
		},
		ValidArgsFunction:     cobra.NoFileCompletions,
		DisableFlagsInUseLine: true,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.Force, "force", "f", false, "Force this node to leave the swarm, ignoring warnings")
	return cmd
}
