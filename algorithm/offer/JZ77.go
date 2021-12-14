package offer

/**
	按之字型打印二叉树
 */

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func printBinaryTree(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	level := 0
	for len(stack) != 0 {
		length := len(stack)
		slic := []int {}
		for i:=0; i < length; i++ {
			node := stack[i]
			slic = append(slic, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[length:]
		if level % 2 != 0 {
			reverseSlice(slic)
		}
		ans = append(ans, slic)
		level++
	}
	return ans
}

func reverseSlice(array []int) {
	left, right := 0, len(array)-1
	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
	return
}
