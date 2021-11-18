package main

import "fmt"

/**
	fmt包中有个Stringer的接口
	Type Stringer interface{ String() string}
	fmt.Println()或fmt.Printf的%v参数 会判断变量是否实现了Stringer接口，若实现，调用String()函数并将返回值打印在屏幕上
 */

func main() {
	fmt.Println(integer(4))
	fmt.Printf("%v", integer(5))
}

type integer int

func (i integer) String() string {
	return "hello"
}


