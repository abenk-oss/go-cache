package cache

import "time"

func (i item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}
