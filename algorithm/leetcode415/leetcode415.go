/**
*   @Author:qixuan ding
*   @Date: 2020/12/30 22:26
**/

package main

import (
	"fmt"
	"strconv"
)

// 思路：和leetcode 2解题思路大致相同， flag标志是否进位

func main() {
	fmt.Println(addStrings("123", "345"))
}

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	s := ""
	flag := 0
	for l1 > 0 || l2 >0 {
		sum := flag
		if l1 > 0 {
			sum += int(num1[l1 -1] - '0')
			l1--
		}
		if l2 > 0 {
			sum += int(num2[l2 -1] - '0')
			l2--
		}
		s = strconv.Itoa(sum%10) + s
		flag = sum / 10
	}
	if flag != 0 {
		s = strconv.Itoa(flag) + s
	}
	return s
}