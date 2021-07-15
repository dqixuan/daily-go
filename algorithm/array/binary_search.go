package array

// 根据leetcode参考手册， 总结二分查找。默认数组升序。

// 二分查找一般操作
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := int(uint(left+right) >> 1)
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return -1
}

func searchFirstEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = int(uint(left+right) >> 1)
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[right] == target {
		return right
	}
	return -1
}

// 找到数组中第一个和目标相等的下标
func searchFirstEqualElementV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = int(uint(left+right) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid] != nums[mid-1] {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

// 查找最后一个与target相等的元素，时间复杂度O(logn)
func searchLastEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = int(uint(left+right) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < right {
			left = mid + 1
		} else {
			if mid == len(nums) - 1 || nums[mid] != nums[mid+1] {
				return mid
			}
			left = mid+1
		}
	}
	return -1
}

// 查找第一个大于等于target的元素，时间复杂度O(logn)
func searchFirstEqualOrGreaterElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = left + (right - left) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

// 查找最后一个小于或等于target的元素, 时间复杂度O(logn)
func searchLastSmallerOrEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = left + (right - left) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}


//func binarySearchV2()
