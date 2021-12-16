package offer

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	root := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val:6},
		Right: &TreeNode{Val:26},
	}
	root = Convert(root)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}
}