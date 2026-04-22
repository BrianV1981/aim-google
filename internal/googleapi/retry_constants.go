package googleapi

import "time"

const (
	// MaxRateLimitRetries is the maximum number of retries on 429 responses.
	// Elevated for autonomous A.I.M. Swarm operations.
	MaxRateLimitRetries = 7
	// RateLimitBaseDelay is the initial delay for rate limit exponential backoff.
	RateLimitBaseDelay = 2 * time.Second
	// Max5xxRetries is the maximum retries for server errors.
	Max5xxRetries = 3
	// ServerErrorRetryDelay is the delay before retrying on 5xx errors.
	ServerErrorRetryDelay = 5 * time.Second
)
