package main

import "fmt"

/**
	堆排序
	基础结构是数组，完全二叉树
 */

func adjustHeap(arr []int, idx, length int) {
	key := arr[idx]
	for k := 2*idx+1; k < length; k=2*k+1 {
		if k+1 < length && arr[k+1] > arr[k] {
			k++
		}
		if arr[k] > key {
			arr[idx] = arr[k]
			idx = k
		} else {
			break
		}
	}
	arr[idx] = key
}

func heapSort(arr []int) {
	length := len(arr)
	for i:= length /2 -1; i>=0; i-- {
		adjustHeap(arr, i, length)
	}
	for i:= length-1; i>0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i)
	}
}

func main() {
	arr := []int{3, 2, 1, 10, 3, 20}
	fmt.Println(arr)  // [3 2 1 10 3 20]
	heapSort(arr)
	fmt.Println(arr) // [1 2 3 3 10 20]
}
