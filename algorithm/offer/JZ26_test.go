package offer

import (
	"fmt"
	"testing"
)

func TestHasSubtree(t *testing.T) {
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	root2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(helpHasSubtree(root1, root2))
}