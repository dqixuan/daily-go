package list

/**
	leetcode 203
	1->2->6->4->5->2->6  val=6
	1->2->4->2->5->2
 */

func removeElement(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			cur.Next = head
			head = head.Next
			cur = cur.Next
			cur.Next = nil
		}
	}
	return dummy.Next
}
