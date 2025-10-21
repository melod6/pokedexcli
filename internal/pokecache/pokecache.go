// Package pokecache handles caching of data from pokeapi
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	newCacheEntry := cacheEntry{time.Now(), val}
	c.cacheEntries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ce, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return ce.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
