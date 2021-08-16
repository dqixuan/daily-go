package Palindrome

import (
	"strconv"
)

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

func oncequickSort(arr []int, left, right int) int {
	key := arr[left]
	for left < right {
		for left < right && arr[right] >= key {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[right] <= key {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = key
	return left
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := oncequickSort(arr, left, right)
	quickSort(arr, left, mid-1)
	quickSort(arr, mid+1, right)
}

