package back_tracking

import "fmt"

// 46 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列

// permute1 方法1：交换法
func permute1(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 {
		ans = append(ans, nums)
		return ans
	}
	backTracing1(nums, 0, &ans)
	return ans
}

func backTracing1(nums []int, idx int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, nums...))
	for i, l := idx, len(nums); i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			backTracing1(nums, i+1, ans)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return
}

// permute  方法2：遍历法
func permute(nums []int) [][]int {
	var (
		arr  []int
		ans  [][]int
		isIn = make([]bool, len(nums))
	)
	helper1(nums, isIn, &arr, &ans)
	return ans
}

func helper1(nums []int, isIn []bool, arr *[]int, ans *[][]int) {
	if len(*arr) == len(nums) {
		*ans = append(*ans, append([]int{}, *arr...))
		return
	}
	for i, l := 0, len(nums); i < l; i++ {
		if isIn[i] == true {
			continue
		}
		*arr = append(*arr, nums[i])
		isIn[i] = true
		helper1(nums, isIn, arr, ans)
		*arr = (*arr)[:len(*arr)-1]
		isIn[i] = false
	}
	return
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	fmt.Println(permute1(nums))
}
