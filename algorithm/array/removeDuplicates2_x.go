package array

/*

给定一个增序排列数组 nums ，你需要在 原地 删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1：

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 你不需要考虑数组中超出新长度后面的元素。

*/

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	idx, count := 0, 0
	for i, v := range nums {
		if i != idx {
			nums[idx] = nums[i]
		} else {
			count++
			idx++
		}
		if v == nums[idx] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			idx++
		}
	}
	return idx
}