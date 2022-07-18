package cash

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func GetCacheConn() *cache.Cache {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	return cache.New(5*time.Minute, 10*time.Minute)
}
