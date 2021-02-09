package leetcode430

/**
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child * Node
}

/**
思路：


	时间复杂度O(n) 不太准确， 空间复杂度O(1)
 */
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	cur := root
	for cur != nil {
		if cur.Child != nil {
			tmp := cur.Child
			for tmp.Next != nil {
				tmp = tmp.Next
			}
			tmp.Next = cur.Next
			if cur.Next != nil {
				cur.Next.Prev = tmp
			}
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil
		}
		cur = cur.Next
	}
	return root
}
