package pokemon

import "time"

const (
	defaultClientTimeout = 10 * time.Second
	defaultBaseURL       = "https://pokeapi.co/api/v2"
)

type Config struct {
	// The URL of the Pokemon API
	BaseURL string
	// The timeout when making HTTP requests to the API
	ClientTimeout time.Duration
	// If you want in-memory caching enabled
	CacheEnabled bool
	// The time-to-live of the cache entries
	CacheTTL time.Duration
}
