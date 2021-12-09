package back_tracking

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

func combine(n int, k int) [][]int {
	var (
		arr []int
		ans [][]int
	)
	backTracing77(n, 1, k, &arr, &ans)
	return ans
}

func backTracing77(n, idx, k int, arr *[]int, ans *[][]int) {
	if len(*arr) == k {
		*ans = append(*ans, append([]int{}, *arr...))
		return
	}
	for i := idx; i <= n; i++ {
		*arr = append(*arr, i)
		backTracing77(n, i+1, k, arr, ans)
		*arr = (*arr)[:len(*arr)-1]
	}
	return
}
