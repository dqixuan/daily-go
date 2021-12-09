package array

import "math"

/**
	2
	3 4
	6 5 7
	4 1 8 3
	找出自定向下的最小路径和，每一步只能移动到下一行相邻的节点
 */

// 方法1  需要和原数组相同大小的辅助空间
func minimumTotal1(triangle [][]int) int {
	rows := len(triangle)
	if rows == 0 {
		return 0
	}
	dp := make([][]int,rows)
	result := math.MaxInt32
	for i:=0; i < rows; i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if i == 0 {
				dp[i][j] += triangle[i][j]
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][0] + triangle[i][j]
				} else if j == len(triangle[i]) - 1 {
					dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
				}
			}
		}
	}
	for i:=0; i < len(dp[rows-1]); i++ {
		result = min(result, dp[rows-1][i])
	}
	return result
}

// 方法2 逆向思维  O(1)空间复杂度
func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	for i := len(triangle)-2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

// 方法3 一维切片的辅助空间
func minimumTotal3(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	ans := make([]int, len(triangle[len(triangle)-1]))
	ans[0] = triangle[0][0]
	for i:=1; i < len(triangle); i++ {
		for j := len(triangle[i])-1; j >= 0; j-- {
			if j == len(triangle[i])-1 {
				ans[j] = triangle[i][j] + ans[j-1]
			} else if j > 0 {
				ans[j] = triangle[i][j] + min(ans[j], ans[j-1])
			} else {
				ans[j] += triangle[i][j]
			}
		}
	}
	result := math.MaxInt32
	for _, v := range ans {
		result = min(result, v)
	}
	return result
}



func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
