package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Mu      sync.RWMutex
	Entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	entries := make(map[string]cacheEntry)
	cache := &Cache{
		Entries: entries,
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	created := time.Now()
	entr := c.Entries

	_, ok := c.Entries[key]
	if !ok {
		entr[key] = cacheEntry{
			createdAt: created,
			val:       val,
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanUp(interval)
	}
}

func (c *Cache) cleanUp(interval time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	current := time.Now()
	for key, entry := range c.Entries {
		if current.Sub(entry.createdAt) > interval {
			delete(c.Entries, key)
		}
	}
}
