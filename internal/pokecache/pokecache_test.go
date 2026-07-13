package pokecache

import (
	"testing"
	"time"
)

var cases = []struct {
	key string
	val []byte
}{
	{
		key: "https://example.com",
		val: []byte("testdata"),
	},
}

func TestAddGet(t *testing.T) {

	for _, c := range cases {
		// Create cache
		cache := NewCache(5 * time.Second)
		// Add c.key and c.val
		cache.Add(c.key, c.val)
		// Get c.key
		cachedVal, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key %q", c.key)
			return
		}

		// Check the result
		if string(cachedVal) != string(c.val) {
			t.Errorf("expected %q, got %q", string(c.val), string(cachedVal))
			return
		}
	}
}
