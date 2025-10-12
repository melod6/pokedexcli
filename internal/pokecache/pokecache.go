package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        sync.Mutex
}

func NewCache(interval time.Duration) {
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheEntries[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if ok := c.cacheEntries[key]; ok == nil {
		return nil, false
	}
	return c.cacheEntries[key], true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.cacheEntries {
			if time.Now().Sub(entry.createdAt) > interval {
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
