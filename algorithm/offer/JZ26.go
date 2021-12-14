package offer

/**
	判断是否是子树结构
 */

func HasSubtree( pRoot1,  pRoot2 *TreeNode ) bool {
	// write code here
	if pRoot2 == nil {
		return false
	}
	return helpHasSubtree(pRoot1, pRoot2)
}

func helpHasSubtree(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil  {
		return false
	}
	if root1.Val == root2.Val && helpHasSubtree(root1.Left, root2.Left) && helpHasSubtree(root1.Right, root2.Right) {
		return true
	}
	return helpHasSubtree(root1.Left, root2) || helpHasSubtree(root1.Right, root2)
}
