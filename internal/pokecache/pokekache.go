package pokecache

import (
	"sync"
	"time"
)

// cacheEntry stores a cached value along with the time it was added.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache stores cached data and protects it from concurrent access.
type Cache struct {
	cache    map[string]cacheEntry // Holds the cached key-value pairs.
	interval time.Duration         // How long entries should stay in the cache.
	mu       sync.Mutex            // Prevents multiple goroutines from modifying the cache at the same time.
}

// NewCache creates and returns a new empty cache.
func NewCache(interval time.Duration) *Cache {
	newMap := map[string]cacheEntry{}
	newCache := Cache{
		cache:    newMap,
		interval: interval,
	}
	go newCache.reapLoop()
	return &newCache
}

// Add inserts a new value into the cache using the given key.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()         // Lock the cache so only one goroutine can modify it at a time.
	defer c.mu.Unlock() // Unlock automatically when the function finishes.

	// Create a cache entry with the current timestamp.
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	// Store the entry in the cache.
	c.cache[key] = entry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		// this block runs once every time interval passes
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}

}
