package offer

/**
	0-9 		9*1
	10-99  		90*2
	100-999		1000*3
 */
func findNthDigit( n int ) int {
	// write code here
	if n < 10 {
		return n
	}
	i := 1

	for !(n >= pow10(i) && n < pow10(i+1)) {
		i++
	}
	return 0
}

func pow10(num int) int {
	ans := 1
	for num > 0 {
		ans *= 10
		num--
	}
	return ans
}
