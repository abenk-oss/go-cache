package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	items map[K]item[V]
	mu    sync.RWMutex
}

type item[V any] struct {
	value  V
	expiry time.Time
}

// New initializes a new Cache instance and launches a goroutine
// that periodically removes expired items from the cache based on the
// specified cleanupInterval.
func New[K comparable, V any](cleanupInterval time.Duration) *Cache[K, V] {

	c := &Cache[K, V]{
		items: make(map[K]item[V]),
	}

	go func() {

		for range time.Tick(cleanupInterval) {

			c.mu.Lock()

			var expiredKeys []K

			for k, item := range c.items {
				if item.isExpired() {
					expiredKeys = append(expiredKeys, k)
				}
			}

			for _, k := range expiredKeys {
				c.delete(k)
			}

			c.mu.Unlock()
		}
	}()

	return c
}

// Set inserts an item to the cache, replacing any existing one.
func (c *Cache[K, V]) Set(key K, data V, ttl time.Duration) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.set(key, data, ttl)
}

// Add inserts an item into the cache if no existing item is associated
// with the given key or if the current item has expired. If an active
// item exists for the key, it returns an error indicating that the item cannot
// be added.
func (c *Cache[K, V]) Add(key K, data V, ttl time.Duration) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.items[key]; found {

		if item.isExpired() {
			c.delete(key)
		} else {
			return fmt.Errorf("item %v already exists", key)
		}
	}

	c.set(key, data, ttl)
	return nil
}

// Replace updates the value for a cache key only if the key already exists
// and the associated item has not expired. If the item has expired, it
// attempts to delete it and returns an error indicating that the value
// cannot be replaced.
func (c *Cache[K, V]) Replace(key K, data V, ttl time.Duration) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if i, found := c.items[key]; found {

		if i.isExpired() {
			c.delete(key)
			return fmt.Errorf("item %v is expired", key)
		} else {
			c.set(key, data, ttl)
			return nil
		}
	}

	return fmt.Errorf("item %v doesn't exist", key)
}

// Get retrieves the value associated with the specified key from the cache.
// It returns the item value along with a boolean indicating whether the key
// was found. If the key is expired, it is deleted from the cache, and the
// function returns false.
func (c *Cache[K, V]) Get(key K) (V, bool) {

	c.mu.Lock()
	defer c.mu.Unlock()

	i, found := c.items[key]
	if !found {
		return i.value, false
	}
	if i.isExpired() {
		c.delete(key)
		return i.value, false
	}

	return i.value, true
}

// Pop deletes and returns the item associated with the specified key from the cache.
// It returns the item value along with a boolean indicating whether the key was found.
// If the key is not found or the item has expired, it deletes the expired item and
// returns the zero value for the item type along with false.
func (c *Cache[K, V]) Pop(key K) (V, bool) {

	c.mu.Lock()
	defer c.mu.Unlock()

	i, found := c.items[key]
	if !found {
		return i.value, false
	}

	c.delete(key)

	if i.isExpired() {
		return i.value, false
	}

	return i.value, true
}

// Remove removes the item associated with the specified key from the cache.
// If the key exists, the item is permanently deleted; if the key is not found,
// no action is taken.
func (c *Cache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.delete(key)
}

// RemoveExpired removes all expired items from the cache.
func (c *Cache[K, V]) RemoveExpired() {

	c.mu.Lock()
	defer c.mu.Unlock()

	var expiredKeys []K

	for key, i := range c.items {
		if i.isExpired() {
			expiredKeys = append(expiredKeys, key)
		}
	}

	for _, key := range expiredKeys {
		c.delete(key)
	}
}

// Clear clears the cache, removing all items.
func (c *Cache[K, V]) Clear() {

	c.mu.Lock()
	defer c.mu.Unlock()

	clear(c.items)
}
