package offer

/**

 */

func Mirror( pRoot *TreeNode ) *TreeNode {
	// write code here
	if pRoot == nil {
		return pRoot
	}
	if pRoot.Left == nil && pRoot.Right == nil {
		return pRoot
	}
	pRoot.Left, pRoot.Right = Mirror(pRoot.Right), Mirror(pRoot.Left)
	return pRoot
}