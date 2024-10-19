# go-cache

A thread-safe, generic in-memory key-value store library for Go, optimized for single-machine applications.

## Installation

To install the package, use the following command:

```bash
go get -u github.com/abenk-oss/go-cache
```

## Development Objectives

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
- [x] **Write Comprehensive Unit Tests for The Implemented Features**

## Usage

```go
package main

import (
 "fmt"
 "time"
 "github.com/abenk-oss/go-cache"
)

func main() {

 // Initialize a new cache with a cleanup interval of 10 seconds.
 // This will automatically remove expired items in the background.
 c := cache.New[string, string](10 * time.Second)

 // 1. Set a key-value pair with a TTL (Time to Live) of 3 seconds.
 c.Set("sessionToken", "abc123", 3*time.Second)
 
 // Retrieve the item before it expires.
 if value, found := c.Get("sessionToken"); found {
  fmt.Println("Session token:", value) 
 }

 // Wait for 4 seconds (allowing the TTL to expire).
 time.Sleep(4 * time.Second)

 // Attempt to retrieve the value again after it has expired.
 if _, found := c.Get("sessionToken"); !found {
  fmt.Println("Session token has expired or not found")  
 }

 // 2. Add an item to the cache only if it doesn't already exist.
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

 // 3. Replace an existing value in the cache if it's still valid.
 err = c.Replace("username", "jane_doe", 10*time.Second)
 if err != nil {
  fmt.Println("Error:", err)
 } else {
  fmt.Println("Replaced username successfully")
 }

 // 4. Clear the cache entirely.
 c.Clear()
 fmt.Println("Cache cleared")
}

```

## Contributing

We welcome contributions to this project!

## License

This project is licensed under the [MIT License](LICENSE).
