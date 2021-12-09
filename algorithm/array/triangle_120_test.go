package array

import (
	"fmt"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	fmt.Println(minimumTotal3([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}