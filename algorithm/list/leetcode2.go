package list

/*
	时间复杂度O(max(m, n))
	空间复杂度O(max(m, n))
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		newHead = &ListNode{}
		cur = newHead
		flag = 0
	)

	for l1 != nil || l2 != nil {
		sum := flag
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		newNode := &ListNode{Val:sum % 10}
		cur.Next = newNode
		cur = cur.Next
		flag = sum / 10
	}
	if flag != 0 {
		cur.Next = &ListNode{Val:flag}
	}
	return newHead.Next
}