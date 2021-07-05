package main

/**
  name:丁其轩
  date:2021/7/5
  time:21:51
*/

// 方法1：每个位置计算其左右侧的最大值。若两者的最小值大于当前值，则表示该点是低洼处。
// 时间空间复杂度 O(n)  O(n)
func trapping1(arr []int) int {
	length := len(arr)
	arr1 := make([]int, length)
	arr2 := make([]int, length)
	for i:=0; i<length; i++ {
		if i > 0 {
			arr1[i] = arr[i]
			if arr1[i] < arr1[i-1] {
				arr1[i] = arr1[i-1]
			}
		} else {
			arr1[i] = arr[i]
		}
	}

	for i:= length-1; i >= 0; i-- {
		if i == length-1 {
			arr2[i] = arr[i]
		} else {
			arr2[i] = arr[i]
			if arr2[i+1] > arr[i] {
				arr2[i] = arr2[i+1]
			}
		}
	}
	ans := 0
	for i := 0; i < length; i++ {
		ans += minValue(arr1[i], arr2[i]) - arr[i]
	}

	return ans
}

func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 方法2：单调递减栈。存储下标。
func trap2(arr []int) int {
	var stack []int
	ans := 0
	for i:=0; i< len(arr); i++ {
		for len(stack) != 0 && arr[i] > arr[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			value := stack[len(stack)-1]
			ans += (i-value-1)*(minValue(arr[value], arr[i])-arr[top])
		}
		stack = append(stack, i)
	}
	return ans
}

// 方法3： 双指针法
func trap(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	left, right := 0, len(arr)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		leftMax = max(leftMax, arr[left])
		rightMax = max(rightMax, arr[right])
		if arr[left] < arr[right] {
			ans += leftMax - arr[left]
			left++
		} else {
			ans += rightMax - arr[right]
			right--
		}
	}

	return ans
}
