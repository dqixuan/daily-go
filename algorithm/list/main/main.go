package main

import (
	"fmt"
	"go.mod/algorithm/list"
)

func main() {
	head := &list.ListNode{
		Val:  1,
		Next: &list.ListNode{
			Val:  2,
			Next: &list.ListNode{
				Val:  3,
				Next: &list.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	ct := list.ReorderList(head)
	for ct != nil {
		fmt.Println(ct.Val)
		ct = ct.Next
	}
}
