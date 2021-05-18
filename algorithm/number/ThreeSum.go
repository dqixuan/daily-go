package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	length := len(nums)
	if length < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i <= length-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // 去重
		}
		target := 0 - nums[i]
		left, right := i+1, length-1
		for left <  right {
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ // 找到相同数字的最后一个
				}
				for left < right && nums[right] == nums[right-1] {
					right-- // 找到相同数字的最后一个
				}
				// 去重
				left++
				right--
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{-2,0,0,2,2}))
}
