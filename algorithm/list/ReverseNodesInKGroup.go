package list

import "container/list"

/**
    @auther: 丁其轩
    @date:   2021/4/9
**/

/**
	leetcode 25: K 个一组翻转链表
	思路：利用双向链表 O(k)
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val:-1}
	cur1 := dummy
	cur2 := head
	doubleList := list.New()
	for cur2 != nil {
		doubleList.PushBack(cur2)
		tmp := cur2
		cur2 = cur2.Next
		tmp.Next = nil
		if doubleList.Len() == k {
			for doubleList.Len() != 0 {
				e := doubleList.Back()
				val := doubleList.Remove(e)
				cur1.Next = val.(*ListNode)
				cur1 = cur1.Next
			}
		}
	}
	for doubleList.Len() != 0 {
		e := doubleList.Front()
		val := doubleList.Remove(e)
		cur1.Next = val.(*ListNode)
		cur1 = cur1.Next
	}
	return dummy.Next
}