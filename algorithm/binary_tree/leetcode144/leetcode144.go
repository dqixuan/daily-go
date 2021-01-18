/**
*   @Author:qixuan ding
*   @Date: 2021/1/18 20:51
**/

package main

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

// 方法1：递归实现 利用闭包原理
func recursiveTraversal(root *TreeNode) (ans []int) {
	cur := root
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(cur)
	return
}

// 方法2：迭代实现， 利用栈实现
func stackTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}

	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			ans = append(ans, cur.Val)
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			cur = cur.Left
		}
		if len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return  ans
}

// 方法3：morris算法
func morrisTraversal(root *TreeNode) []int {
	ans := []int {}
	cur := root

	for cur != nil {
		if  cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				ans = append(ans, cur.Val)
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				cur = cur.Right
			}
		} else {
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}