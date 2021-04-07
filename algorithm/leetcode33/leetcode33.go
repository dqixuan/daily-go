package main

import "fmt"

/**
33. 搜索旋转排序数组
思路：二分查找的变异形式：1，判断中间位置是否是目标；2，寻找递增区间，并判断目标是否在递增区间内
 */

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if target >= nums[left] && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3,5, 1}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}

