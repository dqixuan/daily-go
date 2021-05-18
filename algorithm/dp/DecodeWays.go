package main

/**
    @auther: 丁其轩
    @date:   2021/4/21
    @time:   23:42
**/

/**
91. 解码方法 'A' -> 1  'B' -> 2 ... 'Z' -> 26
思路：
 */

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i:=1; i < len(dp); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

