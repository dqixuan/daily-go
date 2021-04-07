package main

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			if nums[mid] <= nums[right] {
				right = mid-1
			}
		}
	}
	return false
}
