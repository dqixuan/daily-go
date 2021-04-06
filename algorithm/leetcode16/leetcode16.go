package main

import (
	"fmt"
	"math"
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，
使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

 */

func threeSumClosest(nums []int, target int) int {
	var best = math.MaxInt32
	l := len(nums)
	sort.Ints(nums)
	update := func(num int) {
		if abs(num-target) < abs(best-target) {
			best = num
		}
	}
	for i:=0; i < l-2; i++ {
		left, right := i+1, l-1
		for left < right {
			sum := nums[i]+nums[left]+nums[right]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return best
}

func abs(num int) int{
	if num >= 0 {
		return num
	}
	return 0 - num
}

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}
