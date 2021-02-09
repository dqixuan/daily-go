package leetcode112

/**

给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径
这条路径上所有节点值相加等于目标和 targetSum 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// TreeNode 节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法1：递归算法
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 设置终止条件
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	// 进入下层递归
	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}

// 方法2：深度优先遍历
func hasPathSum1(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum, 0)
}

func dfs(root *TreeNode, targetSum, sum int) bool {
	if root == nil  {
		return false
	}
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		return sum == targetSum
	}
	if root.Left != nil && dfs(root.Left, targetSum, sum) {
		return true
	}
	if root.Right != nil && dfs(root.Right, targetSum, sum) {
		return true
	}
	return false
}

