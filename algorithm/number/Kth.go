package main

import (
	"container/heap"
	"fmt"
)

// 求第K大的数


func main() {
	h := []int{23,45,54,56,6,3,1,45,8,78}
	fmt.Println(findKth(h, 5))
}


func findKth(arr []int, k int) int {
	h := &IntHeap{}
	*h = append(*h, arr[0:k]...)
	heap.Init(h)
	for i:=k-1; i < len(arr); i++ {
		if (*h)[0] < arr[i] {
			(*h)[0] = arr[i]
			heap.Init(h)
		}
	}
	return (*h)[0]
}



type IntHeap []int

func (ih IntHeap) Len() int {return len(ih)}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(e interface{}) {
	*ih = append(*ih, e.(int))
}

func (ih *IntHeap) Pop() interface{} {
	old := *ih
	n := ih.Len()
	if n == 0 {
		return 0
	}
	e := old[n-1]
	*ih = old[0:n-1]
	return e
}
