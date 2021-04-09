package list

type ListNode struct {
     Val int
     Next *ListNode
}

/**
leetcode 61, 旋转链表
思路：获取链表长度， k % length 表示旋转的长度
	遍历链表获取新链表表头，并获取链表的最后一个节点tail, 将tail.Next执行原链表的头节点。

 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return  head
	}

	length := getLength(head)
	k = k % length
	if k == 0 {
		return head
	}

	cur := head
	for i:=0; i < length-k-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	cur = newHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return newHead
}

func getLength(head *ListNode) int {
	l := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		l++
	}
	return l
}