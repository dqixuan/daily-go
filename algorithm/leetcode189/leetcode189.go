/**
*   @Author:qixuan ding
*   @Date: 2021/1/3 21:28
**/

package leetcode189

// 空间复杂度O(n) 时间复杂度O(n)
func rotate(nums []int, k int)  {
	l := len(nums)
	k = k % l
	var res []int
	res = append(res, nums[l-k:]...)
	res = append(res, nums[:l-k]...)
	for i, v := range res {
		nums[i] = v
	}
	return
}

//
func rotate1(nums []int, k int)  {
	
}