package lru

import list2 "container/list"

type LRU struct {
	capacity int32
	list     list2.List
	M        map[int32]list2.Element
}
