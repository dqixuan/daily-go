package offer

/**
	搜索二叉树的第k个节点

	思路：morris算法获取前序遍历的数组,
 */

func KthNode(root *TreeNode ,  k int ) int {
	// write code here
	//slic := preOrder(root)
	//if k > len(slic) || k <= 0 {
	//	return -1
	//}
	//return slic[k-1]
	if root == nil && k > 0{
		return -1
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			k--
			if k == 0 {
				return cur.Val
			}
			cur = cur.Right
		} else {
			left := cur.Left
			for left.Right != nil && left.Right != cur {
				left = left.Right
			}
			if left.Right == nil {
				left.Right = cur
				cur = cur.Left
			} else {
				left.Right = nil
				k--
				if k == 0 {
					return cur.Val
				}
				cur = cur.Right
			}
		}
	}
	return -1
}

func preOrder(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			ans = append(ans, cur.Val)
			cur = cur.Right
		} else {
			left := cur.Left
			for left.Right != nil && left.Right != cur {
				left = left.Right
			}
			if left.Right == nil {
				left.Right = cur
				cur = cur.Left
			} else {
				left.Right = nil
				ans = append(ans, cur.Val)
				cur = cur.Right
			}
		}
	}
	return ans
}
