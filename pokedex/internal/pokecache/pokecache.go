package pokecache

import (
	"fmt"
	"time"
)

type CacheEntry struct {
	data       []byte
	created_at time.Time
}

type Cache struct {
	Entries  map[string]CacheEntry
	Interval time.Duration
}

func New(interval time.Duration) *Cache {
	c := &Cache{
		Entries:  make(map[string]CacheEntry),
		Interval: interval,
	}
	go c.reapLoop()
	return c
}
func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.Entries[key]
	return val.data, ok
}

func (c *Cache) Add(val []byte, key string) {
	c.Entries[key] = CacheEntry{
		data:       val,
		created_at: time.Now().UTC(),
	}
}

func (c *Cache) reaploop() {
	for t := range time.Tick(c.Interval) {
		c.reap()
		fmt.Printf("Timer ticked. t==%v\n", t)
	}

}
func (c *Cache) reapLoop() {
	timer := time.NewTicker(c.Interval)

	for range timer.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	current := time.Now().UTC().Add(-c.Interval)
	for k, entry := range c.Entries {
		if entry.created_at.Before(current) {
			delete(c.Entries, k)
			
		}
	}
}
