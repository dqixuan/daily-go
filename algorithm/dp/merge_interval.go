package dp

import "sort"

// 区间合并 56

func merge(interval [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	start, end := interval[0][0], interval[0][1]
	for i:=1; i < len(interval); i++ {
		if interval[i][0] > end {
			ans = append(ans, []int{start, end})
			start, end = interval[i][0], interval[i][1]
		} else {
			if interval[i][1] > end {
				end = interval[i][1]
			}
		}
		if i == len(interval) - 1 {
			ans = append(ans, []int{start, end})
		}
	}
	return ans
}
