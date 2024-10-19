package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCacheSetAndGet(t *testing.T) {

	t.Parallel()

	// New cache with a cleanup interval of 1 second.
	c := New[string, int](1 * time.Second)

	c.Set("key1", 10, 0*time.Second)

	if _, found := c.Get("key1"); found {
		t.Fatal("expected item to be expired and not found")
	}

	c.Set("key2", 10, 5*time.Second)

	if value, found := c.Get("key2"); !found || value != 10 {
		t.Fatalf("expected 10, but got %v, found: %v", value, found)
	}

	// Waiting for the item to expire.
	time.Sleep(6 * time.Second)

	if _, found := c.Get("key2"); found {
		t.Fatal("expected item to be expired and not found")
	}
}

func TestCacheAdd(t *testing.T) {

	t.Parallel()

	c := New[string, int](1 * time.Second)

	err := c.Add("key1", 20, 5*time.Second)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	// should return an error.
	err = c.Add("key1", 30, 5*time.Second)
	if err == nil {
		t.Fatal("expected error for existing item, but got none")
	}

	time.Sleep(6 * time.Second)

	// should succeed because the key is expired.
	err = c.Add("key1", 30, 5*time.Second)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

}

func TestCacheReplace(t *testing.T) {

	t.Parallel()

	c := New[string, int](1 * time.Second)

	// should return an error.
	err := c.Replace("key1", 50, 5*time.Second)
	if err == nil {
		t.Fatal("expected error for non-existent item, but got none")
	}

	c.Set("key1", 40, 5*time.Second)
	// should return no error.
	err = c.Replace("key1", 50, 5*time.Second)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if value, found := c.Get("key1"); !found || value != 50 {
		t.Fatalf("expected 50, but got %v, found: %v", value, found)
	}

	time.Sleep(6 * time.Second)

	// should return an error.
	err = c.Replace("key1", 60, 5*time.Second)
	if err == nil {
		t.Fatal("expected error for expired item, but got none")
	}
}

func TestCachePop(t *testing.T) {

	t.Parallel()

	c := New[string, int](1 * time.Second)

	c.Set("key1", 100, 5*time.Second)

	value, found := c.Pop("key1")
	if !found || value != 100 {
		t.Fatalf("expected to pop 100, but got %v, found: %v", value, found)
	}

	// should be missing.
	_, found = c.Pop("key1")
	if found {
		t.Fatal("expected item to be missing after pop")
	}
}

func TestCacheRemove(t *testing.T) {

	t.Parallel()

	c := New[string, int](1 * time.Second)

	c.Set("key1", 200, 5*time.Second)

	c.Remove("key1")

	if _, found := c.Get("key1"); found {
		t.Fatal("expected item to be removed")
	}
}

func TestCacheClear(t *testing.T) {

	t.Parallel()

	c := New[string, int](1 * time.Second)

	c.Set("key1", 300, 5*time.Second)
	c.Set("key2", 400, 5*time.Second)

	// Clear the cache.
	c.Clear()

	if _, found := c.Get("key1"); found {
		t.Fatal("expected item1 to be cleared")
	}
	if _, found := c.Get("key2"); found {
		t.Fatal("expected item2 to be cleared")
	}
}

func TestCacheRemoveExpired(t *testing.T) {

	t.Parallel()

	c := New[string, int](5 * time.Second)

	c.Set("key1", 500, 1*time.Second)

	c.Set("key2", 600, 10*time.Second)

	// Waiting for the first item to expire.
	time.Sleep(2 * time.Second)

	// Manually removing expired items.
	c.RemoveExpired()

	if _, found := c.Get("key1"); found {
		t.Fatal("expected key1 to be expired and removed")
	}
	if value, found := c.Get("key2"); !found || value != 600 {
		t.Fatalf("expected key2 to be present with value 600, but got %v, found: %v", value, found)
	}
}

func TestCacheConcurrencySafety(t *testing.T) {

	t.Parallel()

	c := New[string, int](1 * time.Second)

	var wg sync.WaitGroup

	// Writer goroutine.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			c.Set(key, i, 5*time.Second)
		}(i)
	}

	// Reader goroutine.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			c.Get(key)
		}(i)
	}

	// Remover goroutine.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			c.Remove(key)
		}(i)
	}

	// Waiting for all goroutines to finish.
	wg.Wait()

	// Verifying that the cache is still in a consistent state.
	for i := 50; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		if _, found := c.Get(key); !found {
			t.Errorf("expected key %s to be found, but it's missing", key)
		}
	}
}
