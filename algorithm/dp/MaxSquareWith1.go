package main

/**
	题意：给定一个矩阵，其中的元素为0或者1，要求找出其中元素全为1的面积最大的正方形
	思路：
		1、边界的值=1的dp[i][j] = 1
		2、转移方程： dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])+1
 */
func findMaxSubSquare(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	if col == 0 {
		return 0
	}
	ans := 0
	dp := make([][]int, row)
	for i:=0; i< row; i++ {
		dp[i] = make([]int, col)
		if matrix[i][0] == 1 {
			dp[i][0] = 1
			ans = max(ans, 1)
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == 1 {
			dp[0][i] = 1
			ans = max(ans, 1)
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	ans := a
	if ans > b {
		ans = b
	}
	if ans > c {
		ans = c
	}
	return ans
}
