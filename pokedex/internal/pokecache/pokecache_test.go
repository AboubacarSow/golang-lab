package pokecache

import (
	"testing"
	"time"
)

func TestAddAndGetCache(t *testing.T) {
	cache := New(45 * time.Second)

	//Preparation
	cache.Add([]byte{12, 54, 65}, "val1")

	result, ok := cache.Get("val1")
	expected := []byte{12, 54, 65}

	if !ok == true {
		t.Errorf("Get cache return %t", ok)
	}

	if len(result) != len(expected) {
		t.Errorf("Expected length:%d . Found actual:%d", len(expected), len(result))
	}
	for i, b := range expected {
		if b != result[i] {
			t.Errorf("Expected byte:%d. Found Actual :%d", b, result[i])
		}
	}
}

func TestReaploop(t *testing.T) {
	interval := 30 * time.Millisecond
	cache := New(interval)

	key := "val1"
	cache.Add([]byte{12, 54, 65}, key)

	time.Sleep(40 * time.Millisecond)

	_, ok := cache.Get(key)
	t.Logf("ok=%v", ok)
	if ok {
		t.Errorf("cache with key:%s should have been reaped", "val1")
	}

}
