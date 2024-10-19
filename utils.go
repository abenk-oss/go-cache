package cache

import "time"

func (i item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}

func (c *Cache[K, V]) set(key K, data V, ttl time.Duration) {
	c.items[key] = item[V]{
		value:  data,
		expiry: time.Now().Add(ttl),
	}
}

func (c *Cache[K, V]) delete(key K) {
	delete(c.items, key)
}
