package inorder_traverse

import (
	"fmt"
	"testing"
)

func Test_inOrderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(inOrderTraversal(root))
}