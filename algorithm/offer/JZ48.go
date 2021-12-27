package offer


func lengthOfLongestSubstring( s string ) int {
	// write code here
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int, len(s))
	idx := -1
	ans := 0
	for i:=0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			for idx < index {
				delete(m, s[idx+1])
				idx++
			}
			m[s[i]] = i
		} else {
			ans = max(ans, i-idx)
			m[s[i]] = i
		}
	}
	return ans
}
