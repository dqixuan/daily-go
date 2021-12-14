package offer

/**

 */


func hasPathSum( root *TreeNode ,  sum int ) bool {
	// write code here
	if root == nil {
		return false
	}
	return helpHasPathSum(root, sum)
}

func helpHasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if sum < 0 {
		return false
	}
	if sum == 0 {
		if root.Right == nil && root.Left == nil {
			return true
		} else {
			return false
		}
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}