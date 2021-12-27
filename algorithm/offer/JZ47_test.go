package offer

import (
	"fmt"
	"testing"
)

func Test_maxValue(t *testing.T) {
	fmt.Println(maxValue([][]int{
		{1,3,1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}