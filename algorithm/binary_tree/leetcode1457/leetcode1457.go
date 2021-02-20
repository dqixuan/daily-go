package main

/**
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree
 */

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路：
	1、验证伪回文，回文形式：1221 或 121， 使用异或， 1^2^2^1 = 0  1^2^1 = 2
		节点数字1~9， _,_,_,_,_,_,_,_,_ 字节代表对应数字， 例如有偶数个2：_,_,_,_,_,_,1,_,_
		n&(n-1) 用于消除数字二级制的最后一位为1的，
		对于最后异或的数字， n&(n-1) == 0, =0 伪回文；!=0 不是

	2、递归获取每每条路径
 */


// 方法1：伪回文判断方式，位运算
func pseudoPalindromicPaths1 (root *TreeNode) (count int) {
	if root == nil {
		return
	}

	var helper func(node *TreeNode, sum int)
	helper = func(root *TreeNode, sum int) {
		sum = sum ^ (1 << root.Val)
		if root.Left == nil && root.Right == nil {
			if sum == 0 || sum & (sum - 1) == 0 {
				count++
			}
			return
		}
		if root.Left != nil {
			helper(root.Left, sum)
		}
		if root .Right != nil {
			helper(root.Right, sum)
		}
		return
	}
	helper(root, 0)
	return count
}

// 方法2：显示超出时间限制
func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		arr []int
		ans [][]int
	)
	helper(root, arr, &ans)

	return len(ans)
}

func helper(root *TreeNode, arr []int, ans *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil {
		if isPalindromic(arr) {
			*ans = append(*ans, arr)
		}
		return
	}
	helper(root.Left, arr, ans)
	helper(root.Right, arr, ans)
	return
}


func isPalindromic(arr []int) bool {
	m := make(map[int]int, len(arr))
	for _, v := range arr {
		m[v]++
	}
	oddCount := 0
	for _, v := range m {
		if v % 2 != 0 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
}


