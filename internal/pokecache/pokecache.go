package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	protect  sync.Mutex
	interval time.Duration
}

func NewCache(intv time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: intv,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.protect.Lock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.entries[key] = entry
	defer c.protect.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.protect.Lock()
	defer c.protect.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.protect.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.protect.Unlock()
	}
}
