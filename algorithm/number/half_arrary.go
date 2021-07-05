package main

/**
	找出数组中出现次数超过一半的数，现在有一个数组，已知一个数出现的次数超过了一半，请用O(n)的复杂度的算法找出这个数。
 */

func halfOfArray(arr []int) int {
	target := 0
	count := 0

	for _, val := range arr {
		if val == target {
			count++
		}
		if count == 0 {
			target = val
		} else {
			count--
		}
	}
	return target
}
