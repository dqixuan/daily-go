package main

import "fmt"

func main() {
	s1 := []int {1, 2, 3}  // s1 len:3 cap:3
	s2 := s1[:2] 			// s2 len:2 cap:3
	// 此时 s1 s2 底层指向同一个数组
	fmt.Println("==============")
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("++++++++++++++")
	s2 = append(s2, 50)  // 50 覆盖底层数组的3
	fmt.Println("==============")
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("++++++++++++++")
	s2 = append(s2, 40) // 发生slice扩容， s2底层指向新的数组
	fmt.Println("==============")
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("++++++++++++++")
	s2[0] =1000 // 修改s2指向新数组的第一个值
	fmt.Println("==============")
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("++++++++++++++")
}
