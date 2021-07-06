package array


// 单调递减栈，实现获取滑动窗口内的最大值。
func window(arr []int, k int) []int {
	stack := []int{}
	result := []int{}

	for i, l := 0, len(arr); i< l; i++ {
		for len(stack) != 0  &&  arr[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
		if i >= k-1 {
			if len(stack) > k {
				stack = stack[1:]
			}
			result = append(result, stack[0])
		}
	}



	return result
}
