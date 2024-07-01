package internal

import "time"

type Cache struct {
	size   int
	values []*CacheItem
}

func NewCache(size int) *Cache {
	return &Cache{
		size: size,
	}
}

func (c *Cache) Get(key string) any {
	for _, ci := range c.values {
		if ci.key == key {
			ci.time = time.Now()
			return ci
		}
	}
	return -1
}

func (c *Cache) Set(key string, value any) {
	for _, ci := range c.values {
		if ci.key == key {
			ci.value = value
			ci.time = time.Now()
			return
		}
	}

	if len(c.values) < c.size {
		c.values = append(c.values, &CacheItem{
			key:   key,
			value: value,
			time:  time.Now(),
		})
	} else {
		oldestIndex := -1
		oldestTime := time.Now()

		for i := range c.values {
			if oldestTime.After(c.values[i].time) {
				oldestTime = c.values[i].time
				oldestIndex = i
			}
		}

		c.values[oldestIndex].value = value
		c.values[oldestIndex].key = key
		c.values[oldestIndex].time = time.Now()
	}
}
