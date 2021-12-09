package main

/**
	202 happy number
	19 -> 1+91=82 -> 64+4=68 -> 36+64=100 -> 1+0+0=1
 */

func isHappy(number int) bool {
	if number == 0 {
		return false
	}
	exist := make(map[int]bool)
	ans := 0
	for number != 1 {
		if exist[number] {
			return false
		}
		exist[number] = true
		for number != 0 {
			ans += number%10
			number /= 10
		}
		number = ans
	}
	return true
}
