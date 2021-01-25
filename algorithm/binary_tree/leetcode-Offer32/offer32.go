/**
*   @Author:qixuan ding
*   @Date: 2021/1/25 22:04
**/

package leetcode_Offer32

import "container/list"

/**
请实现一个函数按照之字形顺序打印二叉树，
即第一行按照从左到右的顺序打印，
第二层按照从右到左的顺序打印，
第三行再按照从左到右的顺序打印，其他行以此类推。
 */

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1  利用原生slice
func levelOrder1(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	level := 0
	for len(stack) > 0 {
		l := len(stack)
		tmp := []int{}
		for l > 0 {
			var node *TreeNode
			if level % 2 == 0 {
				node = stack[0]
				stack = stack[1:]
				if node.Left != nil {
					stack = append(stack, node.Left)
				}
				if node.Right != nil {
					stack = append(stack, node.Right)
				}
				tmp = append(tmp, node.Val)
			} else {
				node = stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				if node.Right != nil {
					stack = append([]*TreeNode{node.Right}, stack...)
				}
				if node.Left != nil {
					stack = append([]*TreeNode{node.Left}, stack...)
				}
				tmp = append(tmp, node.Val)
			}
			l--
		}
		level++
		ans = append(ans, tmp)
	}
	return ans
}


// 方法2：利用container list
func levelOrder2(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	lis := list.New()
	lis.PushBack(root)
	level := 0
	for lis.Len() != 0 {
		length := lis.Len()
		slic := []int{}
		for length > 0 {
			var node *TreeNode
			if level % 2 == 0 {
				node = lis.Remove(lis.Front()).(*TreeNode)
				if node.Left != nil {
					lis.PushBack(node.Left)
				}
				if node.Right != nil {
					lis.PushBack(node.Right)
				}
			} else {
				node = lis.Remove(lis.Back()).(*TreeNode)
				if node.Right != nil {
					lis.PushFront(node.Right)
				}
				if node.Left != nil {
					lis.PushFront(node.Left)
				}
			}
			slic = append(slic, node.Val)
			length--
		}
		level++
		ans = append(ans, slic)
	}
	return ans
}

// 方法3：按层遍历， 偶数行反转
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int {}
	if root == nil {
		return ans
	}
	level := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		l := len(stack)
		tmp := make([]int, l)
		for i:=0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			tmp[i] = node.Val
		}
		if level % 2 != 0 {
			for l, r := 0, len(tmp) - 1; l < r; l, r = l+1, r-1 {
				tmp[l], tmp[r] = tmp[r], tmp[l]
			}
		}
		ans = append(ans, tmp)
		level++
	}
	return ans
}