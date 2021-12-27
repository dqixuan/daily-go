package main

import "fmt"

func main(){
	fmt.Println(match("aba", ".*"))
}

func match(str, pattern string) bool {
	if str == "" || pattern == "" {
		return false
	}
	return matchHandler([]byte(str), 0, []byte(pattern), 0)
}

func matchHandler(s []byte, sIndex int, p []byte, pIndex int) bool {
	sLength := len(s)
	pLength := len(p)
	//s结束，p也结束，匹配成功
	//字符串的所有字符匹配整个模式，所以a与aa不匹配，所以是==
	if sIndex == sLength && pIndex == pLength {
		return true
	}
	//s未结束，p结束 或者 s结束，p未结束，匹配失败
	if (sIndex != sLength && pIndex == pLength) || (sIndex == sLength && pIndex != pLength) {
		return false
	}
	//第二位是*
	if pIndex+1 < pLength && string(p[pIndex+1]) == "*" {
		//第一位匹配
		if (sIndex != sLength && s[sIndex] == p[pIndex]) || (sIndex != sLength && string(p[pIndex]) == ".") {
			return matchHandler(s, sIndex, p, pIndex+2) || //aa匹配a*aa
				matchHandler(s, sIndex+1, p, pIndex+2) || //aa匹配a*a
				matchHandler(s, sIndex+1, p, pIndex) //aa匹配a*
		} else { //第一位不匹配
			return matchHandler(s, sIndex, p, pIndex+2)
		}
	}
	//第二位不是*，一一匹配
	if (sIndex != sLength && s[sIndex] == p[pIndex]) || (sIndex != sLength && string(p[pIndex]) == ".") {
		return matchHandler(s, sIndex+1, p, pIndex+1)
	}
	return false
}
