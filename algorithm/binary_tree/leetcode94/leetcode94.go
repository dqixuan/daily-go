/**
*   @Author:qixuan ding
*   @Date: 2021/1/10 22:30
**/

package leetcode94

// 题目：二叉树的中序遍历

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：递归方法  利用了golang闭包
func inOrderTraversal(root *TreeNode) (res []int) {
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		res = append(res, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return
}

// 方法2:迭代思维，利用栈
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}


