package googleapi

import (
	"context"
	"fmt"

	"google.golang.org/api/youtube/v3"

	"github.com/BrianV1981/aim-google/internal/googleauth"
)

func NewYouTube(ctx context.Context, email string) (*youtube.Service, error) {
	if opts, err := optionsForAccount(ctx, googleauth.ServiceYouTube, email); err != nil {
		return nil, fmt.Errorf("youtube options: %w", err)
	} else if svc, err := youtube.NewService(ctx, opts...); err != nil {
		return nil, fmt.Errorf("create youtube service: %w", err)
	} else {
		return svc, nil
	}
}
