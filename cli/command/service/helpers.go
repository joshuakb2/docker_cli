package service

import (
	"context"
	"io"

	"github.com/joshuakb2/docker_cli/cli/command"
	"github.com/joshuakb2/docker_cli/cli/command/service/progress"
	"github.com/joshuakb2/docker_cli/internal/jsonstream"
)

// WaitOnService waits for the service to converge. It outputs a progress bar,
// if appropriate based on the CLI flags.
func WaitOnService(ctx context.Context, dockerCli command.Cli, serviceID string, quiet bool) error {
	errChan := make(chan error, 1)
	pipeReader, pipeWriter := io.Pipe()

	go func() {
		errChan <- progress.ServiceProgress(ctx, dockerCli.Client(), serviceID, pipeWriter)
	}()

	if quiet {
		go io.Copy(io.Discard, pipeReader)
		return <-errChan
	}

	err := jsonstream.Display(ctx, pipeReader, dockerCli.Out())
	if err == nil {
		err = <-errChan
	}
	return err
}
