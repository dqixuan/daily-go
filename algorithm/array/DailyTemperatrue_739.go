package array

/**
	739
 */

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	dp := []int{}

	for i := len(temperatures)-1; i >= 0; i-- {
		if i == len(temperatures) - 1 {
			dp = append(dp, i)
			ans[i] = 0
		} else {
			for len(dp) != 0 && temperatures[dp[len(dp)-1]] <= temperatures[i] {
				dp = dp[:len(dp)-1]
			}
			if len(dp) != 0 {
				ans[i] = dp[len(dp)-1] - i
			} else {
				ans[i] = 0
			}
			dp = append(dp, i)
		}
	}
	return ans
}