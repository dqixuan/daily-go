package offer

/**
	根据前序和中序遍历构造二叉树
    前序遍历  根节点 左子树  右子树
    中序遍历	 左子树 根节点  右子树
 */

func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
	// write code here
	if len(pre) == 0 || len(vin) == 0 {
		return nil
	}
	return helpConstructBT( vin, pre,0, len(pre)-1, 0, len(pre)-1)
}


func helpConstructBT(pre, vin []int, ps, pe, vs, ve int) *TreeNode {
	if pe < ps || ve < vs {
		return nil
	}
	root := &TreeNode{
		Val:   vin[vs],
	}
	idx := -1
	for i:=ps; i <= pe; i++ {
		if pre[i] == vin[vs] {
			idx = i
			break
		}
	}
	root.Left = helpConstructBT(pre, vin, ps, idx-1, vs+1, vs+1+idx-ps-1)
	root.Right = helpConstructBT(pre, vin, idx+1, pe, idx+1+vs-ps, ve)
	return root
}

