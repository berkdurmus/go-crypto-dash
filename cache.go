package crypto_price_check_dashboard

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value     map[string]float64
	Timestamp time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Get(key string) (map[string]float64, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, exists := c.items[key]
	if !exists || time.Since(item.Timestamp) > 1*time.Minute {
		return nil, false
	}
	return item.Value, true
}

func (c *Cache) Set(key string, value map[string]float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItem{Value: value, Timestamp: time.Now()}
}
