package leetcode85

import (
	"math"
	"sort"
)

func maximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	arr := make([]int, col)
	area := 0
	for _, arr1 := range matrix {
		for i, v := range arr1 {
			if v == 0 {
				arr[i] = 0
			} else {
				arr[i] += 1
			}
		}
		area = max(area, trap(arr))
	}
	return area
}

func trap(arr []int) int {
	l := len(arr)
	if l == 0 {
		return 0
	}
	var stack []int
	area := 0
	idx := 0
	for i, v := range arr {
		for len(stack) != 0 && v < arr[stack[len(stack)-1]] {
			idx = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := 0
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			area = max(area, width*arr[idx])
		}
		stack = append(stack, i)
	}
	for len(stack) != 0 {
		idx = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := 0
		if len(stack) == 0 {
			width = l
		} else {
			width = l - stack[len(stack)-1] - 1
		}
		area = max(area, width*arr[idx])
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minTaps(n int, ranges []int) int {
	internals := [][2]int{}
	// 选取合适区间
	for i, val := range ranges {
		left, right := max(0, i-val), min(n, i+val)
		internals = append(internals, [2]int{left, right})
	}
	sort.Slice(internals, func(i, j int) bool {
		return internals[i][0] < internals[j][0]
	})
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for _, p := range internals {
		if dp[p[0]] == math.MaxInt {
			return -1
		}
		for i := p[0]; i <= p[1]; i++ {
			dp[i] = min(dp[i], dp[i]+1)
		}
	}
	return dp[n]
}
