package leetcode39

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
回溯算法：39、40、46、60、77❎、78、140、257、784、1079
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
*/

func combinationSum(candidates []int, target int) [][]int {
	var (
		sum int
		arr []int
		ans [][]int
	)
	backTracing(candidates, sum, target, &arr, &ans)
	return ans
}

func backTracing(candidates []int, sum, target int, arr *[]int, ans *[][]int) {
	if sum == target {
		tmp := make([]int, len(*arr))
		copy(tmp, *arr)
		*ans = append(*ans, tmp)
		return
	}
	if sum > target {
		return
	}
	for i := 0; i < len(candidates); i++ {
		sum += candidates[i]
		*arr = append(*arr, candidates[i])
		backTracing(candidates, sum, target, arr, ans)
		sum -= candidates[i]
		*arr = (*arr)[:len(*arr)-1]
	}
	return
}

// candidates = [2,3,6,7], target = 7,
// 结果：[[2 2 3] [2 3 2] [3 2 2] [7]]   需要剪枝
