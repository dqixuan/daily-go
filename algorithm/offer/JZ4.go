package offer

/**
    @auther: 丁其轩
    @date:   2021/4/20
    @time:   23:12
**/

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

思路1：选择右上角或是左下角的元素  时间复杂度 O(n+m)
右上角：== 返回true， > 目标元素，横向--； < 目标元素， 纵向++
左下角：== 返回true， > 目标原色，纵向--； < 目标元素,  横向++
 */


func findNumberIn2DArray1(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	i, j := 0, col-1
	for  i < row && j >=0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// 思路2：遍历每一行，每一行有序，使用二分查找  时间复杂度 nlog(m) 或 mlog(n)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i:=0; i < len(matrix); i++ {
		left, right := 0, len(matrix[0])-1
		mid := 0
		for left <= right {
			mid= left + (right - left) >> 1
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}