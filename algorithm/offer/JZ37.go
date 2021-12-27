package offer

import (
	"fmt"
	"strconv"
	"strings"
)

/**
	序列化二叉树
 */


func Serialize( root *TreeNode ) string {
	// write code here
	if root == nil {
		return "#!"
	}
	ans := ""
	ans += fmt.Sprintf("%d!", root.Val)
	ans += Serialize(root.Left)
	ans += Serialize(root.Right)
	return ans
}
func Deserialize( s string ) *TreeNode {
	// write code here
	slic := strings.Split(s, "!")
	if len(slic) == 0 {
		return nil
	}
	return helpDeserialize(&slic)
}

func helpDeserialize(slic *[]string) *TreeNode {
	if len(*slic) == 0 {
		return nil
	}
	val := (*slic)[0]
	if val == "#" {
		*slic = (*slic)[1:]
		return nil
	}
	value, _ := strconv.ParseInt(val, 10, 64)
	root := &TreeNode{Val:int(value)}
	*slic = (*slic)[1:]
	root.Left = helpDeserialize(slic)
	root.Right = helpDeserialize(slic)
	return root
}