package main

import (
	"fmt"
	"sort"
	"time"
)

/**
	LFU:最近不经常使用。
	把数据加载到链表中，按频次排序，一个数据被访问过频次加1， 发生淘汰的时候，把频次低的淘汰掉
	涉及问题：
		1、Get方法O(1)时间复杂度， 选用map
		2、需要对数据按访问频次排序  用slice实现, 首先按frequency降序， 其次按create时间降序
		3、自定义结构体entry
	TODO: 未处理并发
 */

type LFUCache struct {
	capacity 	int
	cacheMap 	map[interface{}]*entry
	list	 	[]*entry
}

type entry struct {
	key 		interface{}
	value 		interface{}
	frequency 	int
	createdAt   time.Time
}

const timeLayout = "2006-01-02 15:04-05"

func (e *entry) String() string {
	return fmt.Sprintf("key:%v, value: %v， frequency:%d time:%s", e.key, e.value, e.frequency, e.createdAt.Format(timeLayout))
}

//
func NewLURCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		cacheMap: make(map[interface{}]*entry),
		list:     make([]*entry, 0),
	}
}

func (lc *LFUCache) sort() {
	sort.Slice(lc.list, func(i, j int) bool {
		if lc.list[i].frequency == lc.list[j].frequency {
			return lc.list[i].createdAt.After(lc.list[j].createdAt)
		}
		return lc.list[i].frequency > lc.list[j].frequency
	})
}

func (lc *LFUCache) Get(key interface{}) (result interface{}, exist bool) {
	val, ok := lc.cacheMap[key]
	if !ok {
		return 
	}
	result = val.value
	val.frequency++
	val.createdAt = time.Now()
	exist = true
	lc.sort()
	return 
}

func (lc *LFUCache) Put(key, value interface{}) {
	val, ok := lc.cacheMap[key]
	if ok {
		val.frequency++
		val.value = value
		val.createdAt = time.Now()
		lc.sort()
		return
	}
	node := &entry{
		key:       key,
		value:     value,
		frequency: 1,
		createdAt: time.Now(),
	}
	lc.cacheMap[key] = node
	if len(lc.list) < lc.capacity {
		lc.list = append(lc.list, node)
	} else {
		lc.list[lc.capacity-1] = node
	}
	lc.sort()
}

func (lc *LFUCache) PrintList() {
	fmt.Println("--------------")
	for _, val := range lc.list {
		fmt.Println(val)
	}
}


func main() {
	lfu := NewLURCache(3)
	lfu.Put(1, 1)
	lfu.PrintList()
	lfu.Put(2, 2)
	lfu.PrintList()
	lfu.Put(2, 21)
	lfu.PrintList()
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 2)
	lfu.PrintList()
	lfu.Put(4, 2)
	lfu.PrintList()
	lfu.Put(5, 2)
	fmt.Println(lfu.Get(1))
	lfu.PrintList()
	lfu.Put(5, 1)
	lfu.PrintList()
	_, ok := lfu.Get(10)
	fmt.Println(ok)
	lfu.Put(5, 5)
	lfu.PrintList()
}





