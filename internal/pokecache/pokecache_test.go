package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cache := NewCache(interval)

	key := "https://pokeapi.co/api/v2/location-area"
	val := []byte("testdata")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected to find key %q", key)
	}
	if string(got) != string(val) {
		t.Fatalf("expected %q, got %q", val, got)
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond

	cache := NewCache(interval)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte("testdata"))

	// Immediately: should exist
	if _, ok := cache.Get("https://pokeapi.co/api/v2/location-area"); !ok {
		t.Fatalf("expected key to be present")
	}

	// Wait a bit longer than interval so reaper can run
	time.Sleep(interval + 5*time.Millisecond)

	if _, ok := cache.Get("https://pokeapi.co/api/v2/location-area"); ok {
		t.Fatalf("expected key to be reaped")
	}
}
