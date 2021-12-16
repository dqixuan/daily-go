package offer

/**
	二叉树转换为双向链表
 */

func Convert( root *TreeNode ) *TreeNode {
	// write code here
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	var head, pre *TreeNode
	cur := root
	for cur != nil {
		if cur.Left == nil {
			if pre == nil {
				head = cur
				pre = cur
				cur = cur.Right
			} else {
				pre.Right = cur
				cur.Left = pre
				pre = pre.Right
				cur = cur.Right
			}
		} else {
			node := cur.Left
			for node.Right != nil && node.Right != cur {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = cur
				cur = cur.Left
			} else {
				if pre == nil {
					head = cur
					pre = cur
					cur = cur.Right
				} else {
					pre.Right = cur
					cur.Left = pre
					pre = pre.Right
					cur = cur.Right
				}
			}
		}
	}
	return head
}
