package array

/**
	给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
	假设只有一个重复的整数，找出这个重复的数。
	e.g [1,3,4,2,2] --> 2
 */

// 方法1： 借助map  O(n)  O(n)
func findDuplicate(arr []int) int {
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
		if m[v] > 1 {
			return v
		}
	}
	return 0
}

func findDuplicate1(arr []int) int {
	return 0
}
