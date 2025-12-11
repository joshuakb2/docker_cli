package client

import (
	"context"

	"github.com/joshuakb2/moby/api/types/registry"
)

// staticAuth creates a privilegeFn from the given registryAuth.
func staticAuth(registryAuth string) registry.RequestAuthConfig {
	return func(ctx context.Context) (string, error) {
		return registryAuth, nil
	}
}
