package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9 
	fmt.Println(twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		if k, ok := m[target - v]; ok {
			return []int {k, i + 1}
		}
		m[v] = i + 1
	}
	return []int {}
}