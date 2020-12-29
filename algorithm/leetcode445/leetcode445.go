package main

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。


示例：
输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// ListNode 链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var s1, s2 []int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			s1 = append(s1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			s2 = append(s2, l2.Val)
			l2 = l2.Next
		}
	}
	var (
		idx1 = len(s1) - 1
		idx2 = len(s2) - 1
		flag = 0
		cur = (*ListNode)(nil)
	)
	for idx1 >= 0 || idx2 >= 0 {
		sum := flag
		if idx1 >= 0 {
			sum += s1[idx1]
			idx1--
		}
		if idx2 >= 0 {
			sum += s2[idx2]
			idx2--
		}
		node := &ListNode{Val:sum%10}
		if cur == nil {
			cur = node
		} else {
			node.Next = cur
			cur = node
		}
		flag = sum / 10
	}
	if flag > 0 {
		node := &ListNode{Val:1}
		node.Next = cur
		cur = node
	}
	return cur
}

/*
思路：输入的链表不能修改，而数字的低位在链表尾部，用栈存储数字，之后处理思路和leetcode2类似
 */
