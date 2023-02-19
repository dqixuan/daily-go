package leetcode42

/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
// 方法1 单调递减栈
func trap(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}
	count := 0
	stack := []int{}
	for i := 0; i < l; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 {
				last := height[stack[len(stack)-1]]
				if last > height[i] {
					stack = append(stack, i)
					break
				} else if last == height[i] {
					stack[len(stack)-1] = i
					break
				}
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					minV := min(height[i], height[stack[len(stack)-1]])
					count += (minV - last) * (i - stack[len(stack)-1] - 1)
				} else {
					stack = append(stack, i)
				}
			}
		}

	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
