package array

/**
	[100, 4, 200, 1, 3, 2]   1, 2, 3, 4   4个数
	要求复杂度 O(n)
 */

func longestConsecutive(array []int) int {
	if len(array) < 2 {
		return len(array)
	}
	m := make(map[int]bool, len(array))
	visit := make(map[int]bool, len(array))
	for _, val := range array {
		m[val] = true
	}
	ans := 0
	for _, v := range array {
		if visit[v] {
			continue
		}
		left, right := v, v
		for m[left-1] {
			left--
			visit[left-1] = true
		}
		for m[right+1] {
			right++
			visit[right+1] = true
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
