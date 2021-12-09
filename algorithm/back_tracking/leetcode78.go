package back_tracking


/**
leetcode 78 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

func subsets(nums []int) [][]int {
	var (
		ans [][]int
		arr []int
	)
	if len(nums) == 0 {
		return ans
	}
	helper78(nums, 0, &arr, &ans)
	return ans
}

func helper78(nums []int, idx int, arr *[]int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, *arr...))

	for i, l := idx, len(nums); i < l; i++ {
		*arr = append(*arr, nums[i])
		helper78(nums, i+1, arr, ans)
		*arr = (*arr)[:len(*arr)-1]
	}
	return
}

