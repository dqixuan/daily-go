package main

import "fmt"

/**
	数字转换为字符串方式：
	1.strconv
	2.fmt.Sprintf("%d", 数字)
	 错误方式： string(数字)， 会把字节或是数字转换为字符的UTF-8的表现形式
 */

func main() {
	i := 12832
	s := string(i)

	fmt.Println(s)
}


