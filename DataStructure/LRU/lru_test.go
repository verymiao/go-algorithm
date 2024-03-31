package LRU

import (
	"reflect"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	tw1 := lru.Get(2)
	lru.Put(2, 4)

	one1 := lru.Get(1)

	if tw1 != 2 {
		t.Logf("lru.Get(2) shoud be 2, but get: %v\n", tw1)
		t.Fail()
	}
	// fmt.Println(one1)

	if !reflect.DeepEqual(one1, -1) {
		t.Logf("lru.Get(1) shoud be -1, but get: %v\n", one1)
		t.Fail()
	}
}
