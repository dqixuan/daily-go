package leetcoce84

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

func largestRectangleArea(heights []int) int {

	return fun1(heights)
}

// 暴力解法 时间复杂度O(n*n) 空间复杂度O(1)
func fun1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	area := 0
	for i := 0; i < len(height); i++ {
		if height[i] == 0 {
			continue
		}
		l, r := i, i
		for l > 0 && height[l-1] >= height[i] {
			l--
		}
		for r < len(height)-1 && height[i] <= height[r+1] {
			r++
		}
		area = max(area, (r-l+1)*height[i])
	}
	return area
}

// 单调栈 存下标
func fun2(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	area := 0
	stack := []int{}
	for i := 0; i < l; i++ {
		for len(stack) != 0 && height[stack[len(stack)-1]] > height[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := 0
			if len(stack) == 0 {
				width = i
			} else {
				width = idx - stack[len(stack)-1] - 1
			}
			area = max(area, width*height[idx])

		}

		stack = append(stack, i)
	}
	for len(stack) != 0 {
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := 0
		if len(stack) == 0 {
			width = l
		} else {
			width = l - stack[len(stack)-1] - 1
		}
		area = max(area, width*height[idx])
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
