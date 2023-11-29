package stack

import (
	list2 "container/list"
	"fmt"
	"testing"
)

func Test_eatFish(t *testing.T) {
	fmt.Println(eatFish([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}))
}

type LRUCache struct {
	cache *list2.List
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: list2.New(),
		value: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key int, value int) {
	//this.cache.PushFront()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
