package main

import (
	"fmt"
	"reflect"
)

/*
	可比较的类型：int float, string, bool pointer, channel, interface, array
	不可比较的类型:map, slice, function
	1、同一结构体比较：是否包含不可比较字段, 若是的话，不可比较；不包含，则可以比较
	2、不同类型结构体：不可直接比较， 如可以强制类型转换， 则转为同一结构体比较

	可以使用reflect.DeepEqual比较
*/

type a struct {
	name string
}

type aa struct {
	name string
	// arr  []int
	m map[int]int
}

type b struct {
	name string
}

func main() {
	// 同一类型结构体， 只包含可比较字段
	a1 := a{name: "a1"}
	a2 := a{name: "a1"}
	fmt.Println("a1 == a2 ：", a1 == a2)

	// 同一类型结构体， 包含不可比较字段
	// aa1 := aa{
	// 	name: "aa1",
	// 	// arr:  []int{12, 3},
	// 	m: map[int]int{1: 1},
	// }
	// aa2 := aa{
	// 	name: "aa1",
	// 	// arr:  []int{12, 3},
	// 	m: map[int]int{1: 1},
	// }
	// fmt.Println("aa1 == aa2 ：", aa1 == aa2)
	// invalid operation: aa1 == aa2 (struct containing []int/map[int]int cannot be compared)

	// 不同结构体比较
	b1 := b{name: "a1"}
	// fmt.Println("a1 == b1: ", a1 == b1)
	// invalid operation: a1 == b1 (mismatched types a and b)

	b2 := (a)(b1)
	fmt.Println("a1 == b2: ", a1 == b2)

	fmt.Println(reflect.DeepEqual(a1, b1))
}
