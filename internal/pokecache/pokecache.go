package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	Entries  map[string]cacheEntry
	interval time.Duration
	lock     sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var result []byte
	c.lock.RLock()
	defer c.lock.RUnlock()
	entry, ok := c.Entries[key]
	if !ok {
		return result, false
	}
	result = entry.value
	return result, true
}

func (c *Cache) reapLoop() {
	t := time.NewTicker(c.interval)
	for clock := range t.C {
		for key, entry := range c.Entries {
			if entry.createdAt.Add(c.interval).Compare(clock) <= 0 {
				delete(c.Entries, key)
			}
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	result := new(Cache)
	result.interval = interval
	go result.reapLoop()
	return result
}
