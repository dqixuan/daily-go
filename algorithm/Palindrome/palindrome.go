package Palindrome

import "strconv"

/**
	回文数  121 true  10 false
	leetcode 9
 */

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}