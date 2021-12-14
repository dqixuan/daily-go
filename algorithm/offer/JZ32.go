package offer

/**
	从上往下打印二叉树
 */

func PrintFromTopToBottom( root *TreeNode ) []int {
	// write code here
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		length := len(stack)
		for i:=0; i < length; i++ {
			ans = append(ans, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[length:]
	}
	return ans
}