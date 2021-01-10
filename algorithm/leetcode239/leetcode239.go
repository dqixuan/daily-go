/**
*   @Author:qixuan ding
*   @Date: 2021/1/2 20:41
**/

package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7}

	rotate(nums, 3)
	fmt.Println(nums)
}
func rotate(nums []int, k int)  {
	l := len(nums)
	k = k % l
	var res []int
	res = append(res, nums[l-k:]...)
	res = append(res, nums[:l-k]...)
	for i, v := range res {
		nums[i] = v
	}
	return
}


func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	l := len(nums)
	if k >= l {
		return []int{nums[maxIndex(nums, 0, l-1)]}
	}
	idx := -1
	var res []int
	for i:= 0;  i < l-k+1; i++ {
		if idx < 0 || idx < i {
			idx = maxIndex(nums, i, i+k-1)
		}
		res = append(res, nums[idx])
		if i + k < l && (nums[i+k] >= nums[idx]) {
			idx = i + k
		}
	}
	return res
}

func maxIndex(nums []int, l, r int) int {
	idx := l
	for i:= l; i <= r; i++ {
		if nums[i] > nums[idx] {
			idx = i
		}
	}
	return idx
}