/**
*   @Author:qixuan ding
*   @Date: 2020/12/29 22:16
**/

package main

import (
	"fmt"
	"strconv"
)

/**
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

输入: num1 = "2", num2 = "3"
输出: "6"

输入: num1 = "123", num2 = "456"
输出: "56088"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// 思路：

func main() {
	s1 := "123"
	s2 := "456"
	fmt.Println(multiply(s1, s2))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1) + len(num2))

	for i:=len(num1) -1; i >=0; i-- {
		a := int(num1[i] - '0')
		for j:=len(num2)-1; j>=0; j-- {
			b := int(num2[j] - '0')
			sum := a * b + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	idx := 0
	for i:=0;i<len(res); i++ {
		if res[i] != 0 {
			idx = i
			break
		}
	}
	ans := ""
	for ; idx < len(res); idx++ {
		ans += strconv.Itoa(res[idx])
	}
	return ans
}
