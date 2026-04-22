package cmd

import (
	"context"

	"github.com/BrianV1981/aim-google/internal/ui"
)

func writeDeleteResult(ctx context.Context, u *ui.UI, resourceName string) error {
	return writeResult(ctx, u,
		kv("deleted", true),
		kv("resource", resourceName),
	)
}
