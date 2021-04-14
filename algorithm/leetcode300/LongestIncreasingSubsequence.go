package main

import "fmt"

/**
    @auther: 丁其轩
    @date:   2021/4/13
**/

// 动态规划 dp[i] 表示以nums[i]结尾的最长递增序列

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	length := len(nums)
	if length == 1 {
		return 1
	}
	dp[0] = 1
	ans := 0
	for i:=1; i< length; i++ {
		tmp := 1
		for j:=0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > tmp {
				tmp = dp[j] + 1
			}
		}
		dp[i] = tmp
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}

// 二分查找法
func method2(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	idx := 0
	for i:=1; i < length; i++ {
		if nums[i] > dp[idx] {
			idx++
			dp[idx] = nums[i]
		} else {
			left, right := 0, idx
			for left < right {
				mid := left + (right-left) >> 1
				if dp[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			dp[left] = nums[i]
		}
	}
	idx++
	return idx
}

func main() {
	arr := []int{10,9,2,5,3,7,101,18}

	fmt.Println(method2(arr))
}
