package main

import (
	"fmt"
	"strings"
)

// 移除无效的括号 1249

func minRemoveToMakeValid(s string) string {
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
func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}
