package offer

/**
	判断是否是平衡二叉树
	平衡二叉树：每个节点的左右子树高度差不超过1
 */


func IsBalanced_Solution( root *TreeNode ) bool {
	// write code here
	if root == nil {
		return true
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if !abs(left, right) {
		return false
	}
	return IsBalanced_Solution(root.Left) && IsBalanced_Solution(root.Right)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func abs(a, b int) bool {
	ans := a - b
	if ans < 0 {
		ans = 0 - ans
	}
	if ans > 1 {
		return false
	}
	return true
}
