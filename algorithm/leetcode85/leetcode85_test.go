package leetcode85

import (
	"fmt"
	"testing"
)

func Test_maximalRectangle(t *testing.T) {
	//data := struct {
	//	A string
	//	B string
	//}{
	//	A: "ding",
	//	B: "b",
	//}
	//
	//b, _ := json.Marshal(data)
	//data1 := struct {
	//	B string
	//	D string
	//}{
	//	B: "b",
	//	D: string(b),
	//}
	//b1, _ := json.Marshal(data1)
	//var obj interface{}
	//err := json.Unmarshal(b1, &obj)
	//fmt.Println(err)
	//fmt.Println(obj)

	//n := 10
	//n += 1 << 3
	//n1 := n + 1<<3
	//fmt.Println(n)
	//fmt.Println(n1)
	//fmt.Println(n >> 3)
	//fmt.Println(n1 >> 3)

	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

}

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func findRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			l = mid
			if l+1 <= r {
				if nums[l+1] == target {
					l++
				} else {
					break
				}
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
