package main

import (
	"math"
	rand2 "math/rand"
	"sync"
	"time"
)

/**
	添加：该节点的层数，找到每一层该节点的插入位置，每一层找的时候都是在前一层的基础上找，而不是在每一层的head重新遍历
	删除：
	查找
 */

const (
	maxLevel = 16
	percentage = 4 // 25%多一层
)

type skipNode struct {
	key int
	value interface{}
	next []*skipNode
}

type skipList struct {
	head, tail *skipNode
	length int
	level int
	m sync.Mutex
	rand *rand2.Rand
}

func NewSkipList(level int) *skipList {
	if level == 0 {
		level = maxLevel
	}
	skipList := &skipList{
		head:   &skipNode{next:make([]*skipNode, level, level)},
		tail:   &skipNode{key:math.MinInt32},
		level:  level,
		rand:   rand2.New(rand2.NewSource(time.Now().UnixNano())),
	}
	for l := range skipList.head.next {
		skipList.head.next[l] = skipList.tail
	}

	return skipList
}

func (sl *skipList) randomLevel() int {
	level := 1
	for level < sl.level && sl.rand.Int() % percentage == 0 {
		level++
	}
	return level
}

func (sl *skipList) Add(key int, value interface{}) {
	// 每一层查找值小于value的最大节点
	sl.m.Lock()
	defer sl.m.Unlock()

	// 新增节点的层数
	level := sl.randomLevel()
	addPos := make([]*skipNode, level)
	node := sl.head
	for index := level-1; index >= 0; index-- {
		for {
			tmp := node.next[index]
			if tmp == sl.tail || tmp.key > key {
				// 下个节点的key大于目标key, 说明当前节点就是当前level的目标节点的前一个节点
				addPos[index] = node
				break
			} else if tmp.key == key {
				// 找到当前节点，修改value
				tmp.value = value
				return
			} else {
				// 替换开始遍历节点
				node = tmp
			}
		}
	}
	nowNode := &skipNode{
		key:   key,
		value: value,
		next:  make([]*skipNode, level),
	}

	for l, pos := range addPos {
		// 执行插入操作
		nowNode.next[l] = pos.next[l]
		pos.next[l] = nowNode
	}
	// 跳表长度加1
	sl.length++
}

func (sl *skipList) Remove(key int) bool {
	sl.m.Lock()
	defer sl.m.Unlock()

	node := sl.head
	removePos := make([]*skipNode, sl.level)
	var target *skipNode
	for level := sl.level-1; level >= 0; level-- {
		for {
			tmp := node.next[level]
			if tmp == sl.tail || tmp.key > key {
				break
			} else if tmp.key == key {
				removePos[level] = node
				target = tmp
				break
			} else {
				node = tmp
			}
		}
	}

	if target != nil {
		for l, pos := range removePos {
			if pos != nil {
				pos.next[l] = target.next[l]
			}
		}
		target = nil
		sl.length--
	}
	return true
}

func (sl *skipList) Search(key int) (interface{}, bool) {
	sl.m.Lock()
	defer sl.m.Unlock()

	node := sl.head
	for level := sl.level-1; level >= 0; level-- {
		for {
			tmp := node.next[level]
			if tmp == sl.tail || tmp.key > key {
				break
			} else if tmp.key == key {
				return tmp.value, true
			} else {
				node = tmp
			}
		}
	}
	return nil, false
}

func (sl *skipList) Len() int {
	sl.m.Lock()
	defer sl.m.Unlock()
	return sl.length
}
