package cache

import (
	"time"
)

type LRUCache struct {
	CacheMap map[string]Cache
	Cap      int
}

type Cache struct {
	value      string
	lastUsed   time.Time
	Expiry     time.Duration
	CreateTime time.Time
}

// Put ...
func (lru *LRUCache) Put(key, value string, expiry time.Duration) bool {
	// Default expiry value
	var leastRecentUsedKey string
	leastRecentUsedValue := time.Now()

	// remove if the max size reached
	if len(lru.CacheMap) >= lru.Cap {
		// remove most recent one
		for key, v := range lru.CacheMap {
			if v.lastUsed.Before(leastRecentUsedValue) {
				leastRecentUsedKey = key
				leastRecentUsedValue = v.lastUsed

			}
		}
		delete(lru.CacheMap, leastRecentUsedKey)
	}

	lru.CacheMap[key] = Cache{
		value:      value,
		lastUsed:   time.Now(),
		CreateTime: time.Now(),
		Expiry:     expiry,
	}
	_, y := lru.CacheMap[key]
	return y
}

// Get ...
func (lru *LRUCache) Get(key string) (string, bool) {
	Cache, prs := lru.CacheMap[key]
	if !prs {
		return "", false
	}

	// Check if the entry has expired
	if time.Since(Cache.CreateTime) > Cache.Expiry {
		delete(lru.CacheMap, key)
		return "", false
	}

	// Update last used time and return the value
	Cache.lastUsed = time.Now()
	lru.CacheMap[key] = Cache
	return Cache.value, true
}
