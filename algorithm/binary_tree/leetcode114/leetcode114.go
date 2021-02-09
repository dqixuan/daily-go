package leetcode114

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路：
	1、当前节点为空，直接返回
	2、当前节点不为空：
		a、左节点为空， 将当前节点指针指向右子树
		b、左节点不为空， 左子树前序遍历的最后一个节点的右孩子指向当前节点的左子树，当前节点右子树为空
		   交换当前节点的左右子树， 将当前节点指针指向右子树

	时间复杂度O(n) 空间复杂度O(1)
 */

func flatten(root *TreeNode)  {
	cur := root

	for cur != nil {
		if cur.Left != nil {
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}

