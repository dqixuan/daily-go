package main

import (
	"math"
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。

思路：先排序，遍历数组， 固定一个数，剩余的两个用双指针法
 */

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	update := math.MaxInt32
	for i:=0; i<=length-2; i++ {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return target
			} else if sum > target {
				right--
			} else {
				left++
			}
			update = updateValue(sum, update, target)
		}
	}
	return update
}

func updateValue(newVal, oldVal, target int) int {
	var abs func(int) int
	abs = func(num int) int {
		if num < 0 {
			return 0 - num
		}
		return num
	}
	if abs(newVal-target) < abs(oldVal-target) {
		return newVal
	}
	return oldVal
}
