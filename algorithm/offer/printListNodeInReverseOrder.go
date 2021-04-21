package offer

type ListNode struct {
	Val int
	Next *ListNode
}

// 题目：逆序打印链表
// 三种方法: 1、递归 O(n) O(n)  2、借助栈 O(n) O(n)  3、先翻转，遍历，再翻转， O(n)  O(1)

// 递归
func printListInReverseOrder(head *ListNode) (ans []int) {
	if head == nil {
		return
	}
	ans = append(ans, printListInReverseOrder(head)...)
	ans = append(ans, head.Val)
	return
}

// 借助栈
func printListInReverseOrder2(head *ListNode) (ans []int) {
	if head == nil {
		return
	}
	cur := head
	for cur != nil {
		ans = append(ans, cur.Val)
		cur = cur.Next
	}
	if len(ans) <= 1 {
		return
	}
	left, right := 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return
}

// 先翻转，遍历，再翻转
func printListInReverseOrder3(head *ListNode) (ans []int) {
	newHead := reverse(head)
	cur := newHead
	for cur != nil {
		ans = append(ans, cur.Val)
		cur = cur.Next
	}
	head = reverse(newHead)
	return
}
// 翻转链表
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	var pre, next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}


