package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	_, err := NewLRUCache(0)
	if err == nil {
		t.Error("wrong")
	}
}

func TestLRUCache2(t *testing.T) {
	lruCache, err := NewLRUCache(2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(lruCache)
	lruCache.Set("1", "a")
	fmt.Println(lruCache)
}
