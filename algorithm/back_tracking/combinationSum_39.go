package back_tracking

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
回溯算法：39❎、40❎、46、60、77❎、78、140、257、784、1079
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
*/

func combinationSum(candidates []int, target int) [][]int {
	var (
		arr []int
		ans [][]int
	)
	helper(candidates, target, arr, &ans)
	return ans
}

func helper(candidates []int, target int, arr []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, arr...))
		return
	}
	if target < 0 {
		return
	}
	for i, l := 0, len(candidates); i < l; i++ {
		arr = append(arr, candidates[i])
		helper(candidates[i:], target-candidates[i], arr, ans)
		arr = arr[:len(arr)-1]
	}
	return
}
