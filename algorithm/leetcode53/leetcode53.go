package main

import "fmt"

// leetcode53 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

/**
思路：数组中数字可能是正、负或是0，
使用的算法应该是贪心算法，
*/

func maxSubArray(arr []int) int {
	// 判断特殊情况
	if len(arr) == 0 {
		return 0
	}

	ans := arr[0] // 不能赋值0， 有可能数组中数字全部小于0
	sum := 0
	for i, l := 0, len(arr); i < l; i++ {
		sum += arr[i]
		if sum > ans {
			ans = sum // 更新返回值
		}
		if sum < 0 {
			sum = 0 // 如果为负数，无论加什么值，只会比当前值小
		}
	}
	return ans
}

func main() {
	num := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(num))
}
