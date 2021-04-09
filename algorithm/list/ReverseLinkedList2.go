package list

/**
    @auther: 丁其轩
    @date:   2021/4/9
**/

/**
	leetcode 92: 反转链表 II
	思路：根据分割位置left, right将链表分为3部分 first, second, third
	first, third都是正常顺序; second部分倒序
	需要记录位置：first部分第一个和最后一个, second部分第一个和最后一个，以及third第一个
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}  // first部分第一个
	cur := dummy
	i := 1

	for ; i < left; i++ {
		cur = cur.Next
	}
	first1 := cur // first部分最后一个
	second1 := cur.Next  // second部分第一个
	cur = second1

	var pre, next *ListNode
	for i <= right {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		i++
	} // pre为second部分第一个， next为third部分第一个

	first1.Next = pre
	second1.Next = next

	return dummy.Next
}