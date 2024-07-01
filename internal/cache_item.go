package internal

import "time"

type CacheItem struct {
	key string
	value any
	time time.Time
}
