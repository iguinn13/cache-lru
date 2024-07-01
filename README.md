# CacheLRU

CacheLRU is a simple implementation of a Least Recently Used (LRU) cache in Go. This project includes a basic cache structure and unit tests to ensure correct functionality.

## Project Structure

The project contains the following files:

- `cache_lru.go`: Implementation of the `Cache` structure and its methods.
- `cache_lru_test.go`: Unit tests for the `Cache` structure.
- `cache_item.go`: Definition of the `CacheItem` structure.
- `cache_item_test.go`: Unit tests for the `CacheItem` structure.

## How to Use

### Installation

Clone the repository to your local machine:

```sh
git clone https://github.com/your-username/CacheLRU.git
```

## Usage

You can use the Cache structure to create an LRU cache of any size. Here is a basic usage example:

```go
package main

import (
    "fmt"
    "time"
    "internal"
)

func main() {
    cache := internal.NewCache(2)
    
    cache.Set("key1", "value1")
    cache.Set("key2", "value2")

    fmt.Println(cache.Get("key1")) // Output: &{key1 value1 <timestamp>}
    fmt.Println(cache.Get("key3")) // Output: -1

    cache.Set("key3", "value3")

    fmt.Println(cache.Get("key1")) // Output: -1 (evicted)
    fmt.Println(cache.Get("key2")) // Output: &{key2 value2 <timestamp>}
    fmt.Println(cache.Get("key3")) // Output: &{key3 value3 <timestamp>}
}

```
