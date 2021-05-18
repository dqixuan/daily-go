package offer

/**
    @auther: 丁其轩
    @date:   2021/4/23
    @time:   22:47
**/

func minArray(numbers []int) int {
	l := len(numbers)
	if l == 0 {
		return 0
	}
	left, right := 0, l-1
	for left < right {
		mid := left + (right - left) >> 1
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return numbers[right]
}