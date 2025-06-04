package pokemon

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Cache is a wrapper around a github.com/patrickmn/go-cache in-memory cache. It is used to cache responses from the
// Pokemon public API to reduce the HTTP traffic and latency.
type Cache struct {
	cache *cache.Cache
}

// NewCache creates a new Cache objects with an in-memory cache with a specified TTL time.Duration.
func NewCache(ttl time.Duration) *Cache {
	return &Cache{cache.New(ttl, ttl)}
}

// GetResponseBodyForURL returns the binary data of the response from the specified URL or nil if it does not exist.
func (c *Cache) GetResponseBodyForURL(url string) []byte {
	entry, isFound := c.cache.Get(url)
	if !isFound {
		return nil
	}

	if data, ok := entry.([]byte); ok {
		return data
	}
	return nil
}

func (c *Cache) CacheResponseForURL(url string, data []byte) {
	c.cache.Set(url, data, 0)
}
