package dp

import "sort"

/**
	72 编辑距离
 */


func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1)
	for i:=0; i < l2; i++ {
		dp[i] = make([]int, l2)
	}
	for i:=0; i < l1; i++ {
		dp[i][0] = i
	}
	for i:=0; i < l2; i++ {
		dp[0][i] = i
	}
	for i:=1; i < l1; i++ {
		for j:=1; j < l2; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minV(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[l1-1][l2-1]
}

func minV(a, b, c int) int {
	ans := a
	if ans > b {
		ans = b
	}
	if ans > c {
		ans = c
	}
	return ans
}


func threeSum(nums []int) [][]int {
	ans := [][]int{}
	length := len(nums)
	if length < 3 {
		return ans
	}
	sort.Ints(nums)
	for i:=0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right :=  i+1, length -1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left< right && nums[left] == nums[left-1] {
					left++
				}
			} else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}