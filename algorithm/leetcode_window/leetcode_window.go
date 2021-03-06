package main

import "fmt"

// 最长无重复子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	m := make(map[byte]int)
	slow := 0
	ans := 0
	for i, l := 0, len(s)-1; i < l; i++ {
		idx, ok := m[s[i]]
		if ok {
			for ; slow <= idx; slow++ {
				delete(m, s[slow])
			}
		}
		m[s[i]] = i
		if ans < i-slow+1 {
			ans = i - slow + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
