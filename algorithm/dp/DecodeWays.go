package dp

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

/**
	自己的理解：

 */
func numDecodings2(s string) int {
	length := len(s)
	dp := make([]int, length)
	for i:=0; i < length; i++ {
		if s[i] - '0' >= 1 && s[i] - '0' <= 9 {
			if i > 0 {
				dp[i] = dp[i-1]
			} else {
				dp[i]=1
			}
		}
		if i > 0 {
			if s[i-1] - '0' > 0 && (s[i-1] - '0') * 10+s[i] - '0' <= 26  {
				if i-2>=0 {
					// "11 23"  中间的两位
					dp[i]+=dp[i-2]
				} else {
					// "11" 最开始的两位
					dp[i] += 1
				}
			}

		}
	}
	return dp[length-1]
}



