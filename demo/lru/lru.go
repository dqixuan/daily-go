package main

import list2 "container/list"

/*
题目：leetcode 146  实现lru算法

思路：
	1、确认要求，lru算法，最少最近使用

*/

type entry struct {
	key   interface{}
	value interface{}
}

type LRUCache struct {
	capacity   int
	doubleList *list2.List
	cache      map[int]*list2.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:   capacity,
		doubleList: list2.New(),
		cache:      make(map[int]*list2.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if ok {
		this.doubleList.MoveToFront(val)
		return val.Value.(*entry).value.(int)
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache == nil {
		this.cache = make(map[int]*list2.Element)
	}
	if v, ok := this.cache[key]; ok {
		this.doubleList.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}
	e := this.doubleList.PushFront(&entry{key, value})
	this.cache[key] = e

	if this.capacity > 0 && this.doubleList.Len() > this.capacity {
		this.removeOldest()
	}
}

func (this *LRUCache) removeOldest() {
	if this.cache == nil {
		return
	}
	e := this.doubleList.Back()
	this.doubleList.Remove(e)
	delete(this.cache, e.Value.(*entry).key.(int))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
