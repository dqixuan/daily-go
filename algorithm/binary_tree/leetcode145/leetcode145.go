/**
*   @Author:qixuan ding
*   @Date: 2021/1/18 21:52
**/

package leetcode145

// 实现二叉树中序遍历

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：递归算法
func recursivePostOrder(root *TreeNode) (ans []int) {
	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		ans = append(ans, root.Val)
	}
	postorder(root)
	return
}

// 方法2:迭代实现
func stackPostOrder(root *TreeNode) ([]int) {
	ans := []int{}
	if root == nil {
		return ans
	}
	s1 := []*TreeNode{root}

	for len(s1) > 0 {
		root := s1[len(s1) - 1]
		s1 = s1[:len(s1) - 1]
		if root.Left != nil {
			s1 = append(s1, root.Left)
		}
		if root.Right != nil {
			s1 = append(s1, root.Right)
		}
		ans = append([]int{root.Val}, ans...)
	}
	return ans
}