package offer

import "sort"

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。 offer 39
/**
	方法1: hashMap 统计 O(n)  	 O(n)
	方法2: 排序         O(nlogn)  O(1)
	方法3: 摩尔投票	   O(n)  	 O(1)
 */

// hashMap
func majorityElement1(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for key, value := range m {
		if value > len(nums)/2 {
			return key
		}
	}
	return -1
}

// 方法2：排序， 取中间值
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 方法3：摩尔投票
func majorityElement(nums []int) int {
	major := -1
	count := 0
	for _, v := range nums {
		if count == 0 {
			major = v
			count++
		} else {
			if v == major {
				count++
			} else {
				count--
			}
		}
	}
	// 确定存在的话， 直接返回major, 下面代码考虑不存在
	counter := 0
	if count <= 0 {
		return -1
	}
	for _, val := range nums {
		if val == major {
			counter++
		}
	}
	if counter > len(nums) / 2 {
		return major
	}
	return -1
}

