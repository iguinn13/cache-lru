package internal

import (
	"testing"
	"time"
)

func TestCacheItemInitialization(t *testing.T) {
	key := "testKey"
	value := "testValue"
	timestamp := time.Now()

	cacheItem := CacheItem{
		key:   key,
		value: value,
		time:  timestamp,
	}

	if cacheItem.key != key {
		t.Errorf("Expected key %s, got %s", key, cacheItem.key)
	}

	if cacheItem.value != value {
		t.Errorf("Expected value %v, got %v", value, cacheItem.value)
	}

	if cacheItem.time != timestamp {
		t.Errorf("Expected timestamp %v, got %v", timestamp, cacheItem.time)
	}
}

func TestCacheItemUpdate(t *testing.T) {
	cacheItem := CacheItem{
		key:   "initialKey",
		value: "initialValue",
		time:  time.Now(),
	}

	newKey := "updatedKey"
	newValue := "updatedValue"
	newTimestamp := time.Now().Add(time.Hour)

	cacheItem.key = newKey
	cacheItem.value = newValue
	cacheItem.time = newTimestamp

	if cacheItem.key != newKey {
		t.Errorf("Expected key %s, got %s", newKey, cacheItem.key)
	}

	if cacheItem.value != newValue {
		t.Errorf("Expected value %v, got %v", newValue, cacheItem.value)
	}

	if cacheItem.time != newTimestamp {
		t.Errorf("Expected timestamp %v, got %v", newTimestamp, cacheItem.time)
	}
}
