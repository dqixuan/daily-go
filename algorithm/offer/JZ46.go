package offer

/**
	 把数字翻译成字符串
	 a->1 b->2 ... z->26
 */

func translate(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	for i:= 0; i < len(s);i++ {
		if s[i] >= '1' && s[i] <= '9' {
			dp[i+1] = dp[i]+1
		}
		if i > 0 && (int(s[i-1]-'0')*10+int(s[i]-'0')) <= 26{
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(s)]
}
