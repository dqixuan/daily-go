package main

import (
	"container/list"
	"fmt"
)

/*
请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func main() {
	root := TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}
	n1 := TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}
	n2 := TreeNode{
		Val: 2,
		Left: nil,
		Right: nil,
	}
	n3 := TreeNode{
		Val: 3,
		Left: nil,
		Right: nil,
	}
	n4 := TreeNode{
		Val: 4,
		Left: nil,
		Right: nil,
	}
	n5 := TreeNode{
		Val: 5,
		Left: nil,
		Right: nil,
	}
	root.Left = &n1
	root.Right = &n2
	n1.Left = &n3
	n1.Right = &n4
	n3.Right = &n5

	// reult := Print(&root)

	r := printTree(&root)
	// fmt.Println(reult)
	fmt.Println(r)
}

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//   采用golang内置结构体 container/list

func printTree(pRoot *TreeNode) [][]int {
	var result [][]int
	if pRoot == nil {
		return result
	}
	l := list.New()
	l.PushBack(pRoot)
	level := 0
	result = append(result, []int{pRoot.Val})
	var node *TreeNode
	for {
		length := l.Len()
		if length == 0 {
			break
		}
		s := []int{}
		if level % 2 == 0 {
			e := l.Back()
			for i := length; i > 0 ;i-- {
				node = e.Value.(*TreeNode)
				if node.Right != nil {
					s = append(s, node.Right.Val)
					l.PushFront(node.Right)
				}
				if node.Left != nil {
					s = append(s, node.Left.Val)
					l.PushFront(node.Left)
				}
				e = e.Prev()
				l.Remove(e.Next())
			}
		} else {
			e := l.Front()
			for i := length; i > 0; i-- {
				node = e.Value.(*TreeNode)
				if node.Left != nil {
					s = append(s, node.Left.Val)
					l.PushBack(node.Left)
				}
				if node.Right != nil {
					s = append(s, node.Right.Val)
					l.PushBack(node.Right)
				}
				
				if e.Next() == nil {
					l.Remove(e)
				} else {
					e = e.Next()
				l.Remove(e.Prev())
				}
			}
		}
		if len(s) != 0 {
			result = append(result, s)
		}
		level++
	}

	return result
}

// Print 打印之字型输出 待优化
func Print( pRoot *TreeNode ) [][]int {
	idx := 0
	var result [][]int
	if pRoot == nil {
		return result
	}
	root := []*TreeNode{pRoot}
	result = append(result, []int{pRoot.Val})
	for {
		var (
			s []int
			ts []*TreeNode
		)
		if idx %2 != 0 {
			for _, v := range root {
				if v.Left != nil {
					ts = append(ts, v.Left)
					s = append(s, v.Left.Val)
				}
				if v.Right != nil {
					ts = append(ts, v.Right) 
					s = append(s, v.Right.Val)
				}
			}
		} else {
			for i:= len(root) - 1; i >= 0; i-- {
				tmp := root[i]
				if tmp.Right != nil {
					// ts = append(ts, tmp.Right) 
					ts = append([]*TreeNode{tmp.Right}, ts...)
					s = append(s, tmp.Right.Val)
				}
				if tmp.Left != nil {
					ts = append([]*TreeNode{tmp.Left}, ts...)
					s = append(s, tmp.Left.Val)
				}
			}
		}
		if len(s) != 0 {
			result = append(result, s)
		}
		if len(ts) == 0 {
			break
		}
		root = ts
		idx++
	}
	return result
}
