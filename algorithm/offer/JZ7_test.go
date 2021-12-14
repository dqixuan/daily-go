package offer

import (
	"fmt"
	"testing"
)

func Test_reConstructBinaryTree(t *testing.T) {
	//[1,2,4,7,3,5,6,8],[4,7,2,1,5,3,8,6]
	pre := []int{1,2,4,7,3,5,6,8}
	vin := []int{4,7,2,1,5,3,8,6}
	root := reConstructBinaryTree(pre, vin)
	fmt.Println(root)
}