package internal

import (
	"testing"
	"time"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := NewCache(2)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	if val := cache.Get("key1"); val.(*CacheItem).value != "value1" {
		t.Errorf("Expected value1, got %v", val)
	}

	if val := cache.Get("key2"); val.(*CacheItem).value != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}

	cache.Set("key3", "value3")

	if val := cache.Get("key1"); val != -1 {
		t.Errorf("Expected -1 for evicted key1, got %v", val)
	}

	if val := cache.Get("key2"); val.(*CacheItem).value != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}

	if val := cache.Get("key3"); val.(*CacheItem).value != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}
}

func TestCacheEviction(t *testing.T) {
	cache := NewCache(2)

	cache.Set("key1", "value1")
	time.Sleep(1 * time.Second)
	cache.Set("key2", "value2")
	time.Sleep(1 * time.Second)
	cache.Set("key3", "value3")

	if val := cache.Get("key1"); val != -1 {
		t.Errorf("Expected -1 for evicted key1, got %v", val)
	}

	if val := cache.Get("key2"); val.(*CacheItem).value != "value2" {
		t.Errorf("Expected value2, got %v", val)
	}

	if val := cache.Get("key3"); val.(*CacheItem).value != "value3" {
		t.Errorf("Expected value3, got %v", val)
	}
}

func TestCacheUpdate(t *testing.T) {
	cache := NewCache(2)

	cache.Set("key1", "value1")
	cache.Set("key1", "newValue1")

	if val := cache.Get("key1"); val.(*CacheItem).value != "newValue1" {
		t.Errorf("Expected newValue1, got %v", val)
	}
}
