package main

import (
	"errors"
	"strconv"
)

var errLocal = errors.New("error is not nil")

func string2Int( s  string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	isNegative := false
	ans := 0
	num := int64(0)
	var err error
	for i:=0; i < len(s); i++ {
		if i == 0 {
			if s[i] == '-' {
				isNegative = true
			} else if s[i] != '+' && !isDigit(s[i]) {
				return 0, errLocal
			} else if  s[i] == '+' {
				continue
			}
			ans = ans * 10 + int(s[i]-'0')
		} else {
			if s[i] == 'e' {
				num, err = strconv.ParseInt(s[i+1:], 10, 64)
				if err != nil {
					return 0, err
				}
				for num > 0 {
					ans *= 10
					num--
				}
				if isNegative {
					return 0-ans, nil
				}
			}
			if !isDigit(s[i]) {
				return 0, errLocal
			}
			ans = ans * 10 + int(s[i]-'0')
		}
	}

	if isNegative {
		ans = 0 - ans
	}
	return ans, nil
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
