package offer

/**
	礼物最大化
 */

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][]int, rows)
	for i:=0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	for i:=0; i < rows; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i:=0; i < cols; i++ {
		if i > 0 {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}
	for i:=1; i < rows; i++ {
		for j:=1; j < cols; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][cols-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}