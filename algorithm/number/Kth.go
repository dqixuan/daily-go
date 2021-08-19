package main

import (
	"container/heap"
	"fmt"
)

// 求第K大的数  1、堆排序  2、变种的快排


func main() {
	h := []int{3,2,3,1,2,4,5,5,6}
	//fmt.Println(findKth(h, 5))
	fmt.Println(getKth(h, 4))
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


// 方法2 变种快排

func getKth(arr []int, k int) int {
	// 默认  len(arr) >=1    0<k<=len(arr)

	return quickSort(arr, 0, len(arr)-1, len(arr)-k)
}

func quickSort(arr []int, left, right, k int) int {

	sentinel := onceQuickSort(arr, left, right)
	if sentinel == k {
		return arr[sentinel]
	} else if sentinel > k {
		return quickSort(arr, left, sentinel-1, k)
	} else {
		return quickSort(arr, sentinel+1, right, k)
	}
}

// 递减排序
func onceQuickSort(arr []int, left, right int) int {
	key := arr[left]
	for left < right {
		for left < right && arr[right] >= key {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= key {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = key
	return left
}

