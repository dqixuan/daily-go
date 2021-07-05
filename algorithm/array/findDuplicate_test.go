package array

import (
	"fmt"
	"testing"
)

func Test_findDuplicate1(t *testing.T) {
	arr := []int{1,3, 4, 2, 2}
	fmt.Println(findDuplicate(arr))
}