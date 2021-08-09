package Palindrome

func matches(num int) []string {
	ans := new([]string)
	*ans = make([]string, 0)
	if num == 0 {
		return *ans
	}

	var help func(string, int, int, int)
	help = func(base string, left, right, num int) {
		if num == left && right == num {
			*ans = append(*ans, base)
			return
		}
		if left == 0 {
			help(base+"(", left+1, right, num)
		} else {
			if left == num {
				help(base+")", left, right+1, num)
			} else {
				base += "("
				left++
				help(base, left, right, num)
				left--
				base = base[:len(base)-1]
				help(base+")", left, right+1, num)
			}
		}

	}

	help("", 0, 0, num)

	return *ans
}

//func help(base string, left, right, num int, ans []string) []string {
//	if left == num && right == num {
//		ans = append(ans, base)
//		return ans
//	}
//	if left < num {
//		return help(base+"(", left+1, right, num, ans)
//	}
//	if right < num && left < right {
//		return help(base+")", left, right+1, num, ans)
//	}
//	return
//}

