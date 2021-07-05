package main

import (
	"fmt"
	"testing"
)

func Test_halfOfArray(t *testing.T) {
	arr := []int{1,2,3,5,4,8,4,4,4}
	fmt.Println(halfOfArray(arr))
}