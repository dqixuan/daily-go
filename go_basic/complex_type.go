package main

import (
	"crypto/sha256"
	"fmt"
)

/**
	数组、切片、map、结构体
	var array [3]int  数组长度不同，数组类型不同
 */
func main() {
	var array [3]int
	fmt.Println(array[0])
	fmt.Println(array[len(array)-1])

	for i, v := range array {
		fmt.Printf("key=%d  value=%d\n", i, v)
	}
	for _, v := range array {
		fmt.Printf("value=%d\n", v)
	}
	var b [10]int = [10]int{1:2,3:3, 9:200, 4:20}
	fmt.Println(b)

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1==c2, c1)
}