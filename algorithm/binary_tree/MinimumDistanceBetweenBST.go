package binary_tree

import "math"

/**
    @auther: 丁其轩
    @date:   2021/4/16
    @time:   22:39
**/

/*
	783. 二叉搜索树节点最小距离
 */

// morris算法  时间O(n) 空间O(1)
func minDiffInBST(root *TreeNode) int {
	cur := root
	var pre *TreeNode
	ans := math.MaxInt32
	helper := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for cur != nil {
		if cur.Left == nil {
			if pre != nil && cur != nil {
				ans = helper(ans, cur.Val - pre.Val)
			}
			pre = cur
			cur = cur.Right
		} else {
			left := cur.Left
			for left.Right != nil && left.Right != cur {
				left = left.Right
			}
			if left.Right == nil {
				left.Right = cur
				cur = cur.Left
			} else {
				left.Right = nil
				if pre != nil && cur != nil {
					ans = helper(ans, cur.Val - pre.Val)
				}
				pre = cur
				cur = cur.Right
			}
		}
	}
	return  ans
}
