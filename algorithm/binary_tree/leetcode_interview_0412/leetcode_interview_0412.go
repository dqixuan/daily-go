package main

import "fmt"

/**
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。
设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。
注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/paths-with-sum-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var num int
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	path(root, sum, 0)
	return num
}

func path(root *TreeNode, sum, count int) {
	if root == nil {
		return
	}
	count += root.Val
	if count == sum {
		num++
	}
	if root.Left != nil {
		path(root.Left, sum, count)
		path(root.Left, sum, 0)
	}
	if root.Right != nil {
		path(root.Right, sum, count)
		path(root.Right, sum, 0)
	}
	return
}

func main() {
	ans := pathSum(nil, 1)
	fmt.Println(ans)
}
