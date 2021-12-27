package offer
/**
	按层打印二叉树
 */

func Print( pRoot *TreeNode ) [][]int {
	// write code here
	ans := [][]int{}
	if pRoot == nil {
		return ans
	}
	stack := []*TreeNode{pRoot}

	for len(stack) != 0 {
		length := len(stack)
		tmp := []int{}
		for i:=0; i<length; i++ {
			tmp = append(tmp, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[length:]
		ans = append(ans, tmp)
	}
	return ans
}
