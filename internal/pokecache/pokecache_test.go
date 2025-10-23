package pokecache

import (
	_ "embed"
	"fmt"
	"testing"
	"time"
)

//go:embed testdata/expected_zone_1_res.json
var expectedJSONzone1 []byte

//go:embed testdata/expected_zone_2_res.json
var expectedJSONzone2 []byte

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/1/",
			val: []byte(expectedJSONzone1),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/2/",
			val: []byte(expectedJSONzone2),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area/1/", []byte(expectedJSONzone1))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area/1/")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area/1/")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
