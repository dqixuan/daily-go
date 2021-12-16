package main

import (

)

/**
    @auther: 丁其轩
    @date:   2021/4/12
**/

func largestNumber(nums []int) string {
	return ""
}

func helper(nums int) int {
	for nums / 10 > 0 {
		nums /= 10
	}
	return nums
}