/**
*   @Author:qixuan ding
*   @Date: 2021/1/10 21:55
**/

package leetcode104

// leetcode-Offer32 104 求二叉树的最大高度

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：使用递归方式
/**
	递归模板：
		1、递归终止条件
		2、处理当前逻辑
		3、进入下层逻辑
		4、清除当前逻辑条件
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if r > l {
		return r + 1
	} else {
		return l + 1
	}
}

// 方法2： 按层遍历二叉树

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	slic := []*TreeNode{root}

	for {
		l := len(slic)
		if l == 0 {
			break
		}
		for l > 0 {
			tmp := slic[0]
			slic = slic[1:]
			if tmp.Left != nil {
				slic = append(slic, tmp.Left)
			}
			if tmp.Right != nil {
				slic = append(slic, tmp.Right)
			}
			l--
		}
		ans++
	}
	return ans
}

