package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 2,4,5}, 3))
}

func reverse(tree []int ) []int {
	tmp := make([]int, len(tree))
	
}

func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	mid := 0
	for left < right {
		mid = left + (right-left) / 2
		if array[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if array[right] == target {
		return right
	}
	return -1
}





