package commands

import (
	"github.com/joshuakb2/docker_cli/cli/command"
	_ "github.com/joshuakb2/docker_cli/cli/command/builder"
	_ "github.com/joshuakb2/docker_cli/cli/command/checkpoint"
	_ "github.com/joshuakb2/docker_cli/cli/command/config"
	_ "github.com/joshuakb2/docker_cli/cli/command/container"
	_ "github.com/joshuakb2/docker_cli/cli/command/context"
	_ "github.com/joshuakb2/docker_cli/cli/command/image"
	_ "github.com/joshuakb2/docker_cli/cli/command/manifest"
	_ "github.com/joshuakb2/docker_cli/cli/command/network"
	_ "github.com/joshuakb2/docker_cli/cli/command/node"
	_ "github.com/joshuakb2/docker_cli/cli/command/plugin"
	_ "github.com/joshuakb2/docker_cli/cli/command/registry"
	_ "github.com/joshuakb2/docker_cli/cli/command/secret"
	_ "github.com/joshuakb2/docker_cli/cli/command/service"
	_ "github.com/joshuakb2/docker_cli/cli/command/stack"
	_ "github.com/joshuakb2/docker_cli/cli/command/swarm"
	_ "github.com/joshuakb2/docker_cli/cli/command/system"
	_ "github.com/joshuakb2/docker_cli/cli/command/volume"
	"github.com/joshuakb2/docker_cli/internal/commands"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, dockerCLI command.Cli) {
	for _, c := range commands.Commands() {
		cmd.AddCommand(c(dockerCLI))
	}
}
