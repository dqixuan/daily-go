package main

import (
	"fmt"
	"strings"
)

// 移除无效的括号 1249

// minRemoveToMakeValid1 记录
func minRemoveToMakeValid1(s string) string {
	if s == "" {
		return ""
	}
	ans := strings.Builder{}
	m := make(map[int]int, 0)
	s1 := []int{}
	for i, l := 0, len(s); i < l; i++ {
		if s[i] == '(' {
			s1 = append(s1, i)
		}
		if s[i] == ')' {
			if len(s1) == 0 {
				m[i] = i
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
	}
	for _, v := range s1 {
		m[v] = v
	}

	for i, l := 0, len(s); i < l; i++ {
		if _, ok := m[i]; !ok {
			ans.WriteByte(s[i])
		}
	}

	return ans.String()
}

// main 待优化

func minRemoveToMakeValid(s string) string {
	if s == "" {
		return ""
	}
	ans := make([]byte, 0)
	balance := 0

	for i, l := 0, len(s); i < l; i++ {
		if s[i] == ')' {
			if balance <= 0 {
				continue
			} else {
				balance--
				ans = append(ans, s[i])
			}
		} else if s[i] == '(' {
			balance++
			ans = append(ans, s[i])
		} else {
			ans = append(ans, s[i])
		}
	}
	// 不能正向遍历，若存在'('，执行if后 i不变，ans[i]指向的内容有变化
	for i := len(ans) - 1; i >= 0 && balance > 0; i-- {
		if ans[i] == '(' {
			ans = append(ans[:i], ans[i+1:]...)
			balance--
		}
	}
	return string(ans)
}
func main() {
	fmt.Println(minRemoveToMakeValid("))(("))
}
