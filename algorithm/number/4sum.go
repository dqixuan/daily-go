package main

import "sort"

/**
	18. 四数之和
 */

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	length := len(nums)
	if length < 4 {
		return ans
	}
	sort.Ints(nums)
	for i:=0; i < length-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i+1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, length-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return ans
}
