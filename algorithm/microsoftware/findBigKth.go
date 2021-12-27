package microsoftware

import (
	"container/heap"
)



func findKth(array []int, n int, K int) int {
	if len(array) == 0 || K == 0 {
		return 0
	}
	return getKth(array, 0, len(array)-1, K)
}

func getKth(array []int, left, right, k int) int {
	idx := onceQuickSort(array, left, right)
	if idx == k-1 {
		return array[idx]
	} else if idx > k+1 {
		return getKth(array, left, idx-1, k)
	} else {
		return getKth(array, idx+1, right, k)
	}
}

func onceQuickSort(array []int, left, right int) int {
	key := array[left]
	l, r := left, right
	for l < r {
		for l<r && array[r] <= key {
			r--
		}
		array[l] = array[r]
		for l<r && array[l] >= key {
			l++
		}
		array[r] = array[l]
	}
	array[l] = key
	return l
}


// 方法1 堆排序
func findKth1( array []int ,  K int ) int {
	// write code here
	if len(array) == 0 || K == 0 {
		return 0
	}
	ha := &heapArray{}
	for i, v := range array {
		if i < K {
			ha.Push(v)
		} else {
			heap.Init(ha)
			if v > (*ha)[0] {
				(*ha)[0] = v
				heap.Init(ha)
			}
		}
	}
	return (*ha)[0]
}

// 堆排序
type heapArray []int

func (ha heapArray) Len() int {
	return len(ha)
}

func (ha heapArray) Swap(i, j int) {
	ha[i], ha[j] = ha[j], ha[i]
}

func (ha heapArray) Less(i, j int) bool {
	return  ha[i] < ha[j]
}

func (ha *heapArray) Push(e interface{}) {
	ele := e.(int)
	*ha = append(*ha, ele)
}

func (ha *heapArray) Pop() interface{} {
	if len(*ha) == 0 {
		return nil
	}
	e := (*ha)[len(*ha)-1]
	*ha = (*ha)[:len(*ha)-1]
	return e
}