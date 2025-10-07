package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cacheEntries     map[string]cacheEntry
	mutex            sync.Mutex
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

func (c *Cache) reapLoop()

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


