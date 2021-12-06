package list

/**
	a1->a2->a3->a4->a5
	a1->a5->a2->a4->a3
	leetcode 143
	思路：将链表分为两部分， 第二部分倒序， 合并链表
 */

func ReorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return  head
	}
	l1, l2 := divideList(head)
	l2 = reverseList(l2)
	return merge(l1, l2)
}


// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//
func divideList(head *ListNode) (l1, l2 *ListNode) {
	l1 = head
	if head == nil || head.Next == nil {
		l2 = nil
		return
	}
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == nil {
		l2 = pre.Next
		pre.Next = nil
		return
	}
	l2 = slow.Next
	slow.Next =  nil
	return
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	i := 0
	for l1 != nil || l2 != nil {
		if i % 2 == 0 {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
			cur.Next = nil
		} else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
			cur.Next = nil
		}
		i++
	}
	return head.Next
}


