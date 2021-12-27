package offer


func minNumberInRotateArray( rotateArray []int ) int {
	// write code here
	l, r := 0, len(rotateArray)-1
	mid := 0

	for l < r {
		mid = l + (r-l) >> 1
		if rotateArray[r] > rotateArray[mid] {
			r = mid
		} else if rotateArray[r] < rotateArray[mid] {
			l = mid+1
		} else {
			l--
		}
	}
	return rotateArray[l]
}