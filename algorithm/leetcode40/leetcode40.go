package leetcode40

import "sort"

/**
title：组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-ii
*/

func combinationSum2(candidates []int, target int) [][]int {
	var (
		arr []int
		ans [][]int
	)
	sort.Ints(candidates)
	backTracing(candidates, 0, target, arr, &ans)
	return ans
}

func backTracing(candidates []int, idx, target int, arr []int, ans *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*ans = append(*ans, append([]int{}, arr...))
		return
	}
	for i, l := idx, len(candidates); i < l; i++ {
		// 剔除相同开始的数字
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}
		arr = append(arr, candidates[i])
		backTracing(candidates, i+1, target-candidates[i], arr, ans)
		arr = arr[:len(arr)-1]
	}
	return
}
