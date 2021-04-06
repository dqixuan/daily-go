package leetcode15

import "sort"

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
 */

func threeSum(nums []int) [][]int {
	l := len(nums)
	var ans [][]int
	if l < 3 {
		return ans
	}
	// 排序
	sort.Ints(nums)
	for i:=0; i < l-2; i++ {
		// 排序后，第一个值相同的忽略
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		right := l-1
		target := 0- nums[i]
		for left:= i+1; left < l; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			for left < right && nums[left] + nums[right] > target {
				right--
			}
			if left == right {
				break
			}
			if nums[left] + nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
			}
		}
	}
	return ans
}


