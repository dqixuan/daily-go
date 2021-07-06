package dp

// 最长有效括号长度
/**
	思路：
		1、当前为'(' dp[i] = 0
		2、当前为')',
			a) i-1位置为'(', dp[i]至少为 2, 若 i-2>=0, dp[i] += dp[i-2]
			b) i-1位置为')', i-dp[i-1]-1位置

 */
func longestValidParentheses(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 0
	result := 0
	for i:=1; i < length; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] += dp[i-2]
				}
				if dp[i] > result {
					result = dp[i]
				}
			} else {
				if tmp:=i - dp[i-1] - 1; tmp >= 0 && s[tmp] == '(' {
					dp[i] = dp[i-1] + 2
					if tmp - 1 >= 0 {
						dp[i] += dp[tmp-1]
					}
					if dp[i] > result {
						result = dp[i]
					}
				}
			}
		}
	}
	return result
}
