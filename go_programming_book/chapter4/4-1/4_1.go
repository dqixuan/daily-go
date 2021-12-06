package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	
	c1 := sha256.Sum256([]byte{'X'})
	fmt.Println(c1)
	fmt.Println(countByte(c1))
}

func countByte(s [32]byte) int {
	ans := 0
	m := make(map[byte]bool, len(s))
	for i:= range s {
		if m[s[i]] {
			continue
		}
		ans++
		m[s[i]]=true
	}
	return ans
}