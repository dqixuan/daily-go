package array

import "fmt"

/*

给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2], 

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。 

你不需要考虑数组中超出新长度后面的元素。


*/

func removeDuplicates(numbers []int) int {
	length := len(numbers)
	if length== 0 || length == 1 {
		return length
	}
	var (
		pre = 0
		cur = 1
	)
	for cur < length {
		if numbers[pre] == numbers[cur] {
			cur++
		} else {
			if pre +1 != cur {
				numbers[pre+1] = numbers[cur]
			}
			pre++
			cur++
		}
	}
	fmt.Println(numbers)
	return pre+1
}

func main1() {
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}