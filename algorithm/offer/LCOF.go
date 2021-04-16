package offer

/**
    @auther: 丁其轩
    @date:   2021/4/16
**/

/**
	找出第一个重复的数字  n个0~n-1的数字中有重复的
 */

// 空间复杂度O(1) 时间O(n)
func findRepeatNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range m {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = 1
	}
	return 0
}

// 空间复杂度O(1) 时间O(n)
func findRepeatNumber1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}
