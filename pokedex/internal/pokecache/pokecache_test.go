package pokecache

import (
	"testing"
	"time"
)

func TestAddCache(t *testing.T) {
	cache := New(45 * time.Second)

	//Preparation
	cache.Add([]byte{}, "val1")
}