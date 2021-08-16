package Palindrome

func matches(num int) []string {
	ans := make([]string, 0)
	if num == 0 {
		return ans
	}

	var help func(string, int, int)
	help = func(base string, left, right int) {
		if num * 2 == len(base) {
			ans = append(ans, base)
			return
		}
		if left > 0 {
			help(base+"(", left-1, right)
		}
		if right > 0 && right > left {
			help(base+")", left, right-1)
		}
	}
	help("", num, num)
	return ans
}



