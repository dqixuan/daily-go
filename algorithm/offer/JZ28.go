package offer

/**
	判断二叉树是否是对称的二叉树
 */

func isSymmetrical( pRoot *TreeNode ) bool {
	// write code here
	if pRoot == nil {
		return true
	}
	return isSymmetric(pRoot.Left, pRoot.Right)
}

func isSymmetric(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return isSymmetric(root1.Left, root2.Right) && isSymmetric(root1.Right, root2.Left)
	}
	return false
}

