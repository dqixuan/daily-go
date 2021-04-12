package main

import "fmt"

/**
    @auther: 丁其轩
    @date:   2021/4/12
**/

/**
	264. 丑数 II
	思路：三指针法
 */

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	nums := []int{2,3,5}
	idx1, idx2, idx3 := 1, 1, 1 // 初始化指针，都指向第一个
	ans := make([]int, n+1)
	ans[1] = 1 // 初始化， 第一个丑数是1
	tmp1, tmp2, tmp3 := 0, 0, 0 // 指定临时变量
	for i:=2; i<=n; i++ {
		// 数字X对应位置的值，选择最小的那个
		tmp1 = nums[0] * ans[idx1]
		tmp2 = nums[1] * ans[idx2]
		tmp3 = nums[2] * ans[idx3]
		ans[i] = min(tmp1, tmp2, tmp3)
		// 如果等于最小的那个， 对应下标+1
		if tmp1 == ans[i] {
			idx1++
		}
		if tmp2 == ans[i] {
			idx2++
		}
		if tmp3 == ans[i] {
			idx3++
		}
	}
	return ans[n]
}

func min(a, b, c int) int {
	ans := a
	if ans > b {
		ans = b
	}
	if ans > c {
		ans = c
	}
	return ans
}

func main() {
	fmt.Println(nthUglyNumber(11))
}
