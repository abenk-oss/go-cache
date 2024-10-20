# go-cache

A thread-safe, generic in-memory key-value cache for Go, ideal for single-machine applications. The cache's key advantage is its use of generics, providing type safety and eliminating the need for casting or working with empty interfaces. Additionally, since it operates entirely in-memory with built-in expiration, there's no need for serialization or network transmission, which improves performance for local caching needs.

## Installation

To install the package, use the following command:

```bash
go get -u github.com/abenk-oss/go-cache
```

<!-- ## Development Objectives

- [x] **Set Up Project Boilerplate**
- [x] **Utilize Go Generics**
- [x] **Implement Core Methods**
  - `Set`: Add or update a value in the cache.
  - `Add`: Insert a value only if no existing value is associated with the key or if the existing item has expired.
  - `Replace`: Only update the value for an existing active key.
  - `Get`: Retrieve an active value by its key, returning an indication of its existence.
  - `Pop`: Remove and return an active value associated with a key.
  - `Remove`: Permanently remove a value associated with a specified key from the cache.
  - `RemoveExpired`: Permanently Remove all expired items from the cache.
  - `Clear`: Empty the cache.
- [x] **Ensure Thread Safety**
- [x] **Write Comprehensive Unit Tests for The Implemented Features** -->

## Usage

```go
package main

import (
 "fmt"
 "time"
 "github.com/abenk-oss/go-cache"
)

func main() {

  // Create a new cache instance with a cleanup interval of 10 seconds.
  // In this example, The cache stores string keys and string items.
  c := cache.New[string, string](10 * time.Second)


  // Set the item of the key "foo" to "bar" with a TTL (Time to Live) of 3 seconds.
  c.Set("foo", "bar", 3*time.Second)

  // Get the string associated with the key "foo" from the cache
  if foo, found := c.Get("foo"); found {
    fmt.Println(foo) 
  }

  // Add an item to the cache only if it doesn't already exist.
  err := c.Add("username", "john_doe", 10*time.Second)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Added username successfully")
  }

  // Attempting to add the same key again will result in an error.
  err = c.Add("username", "jane_doe", 10*time.Second)
  if err != nil {
    fmt.Println("Error:", err)  // Output: Error: item username already exists
  }

  // Clear the cache entirely.
  c.Clear()
}
```

## Contributing

We welcome contributions to this project! For detailed guidelines on how to contribute, please see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).
