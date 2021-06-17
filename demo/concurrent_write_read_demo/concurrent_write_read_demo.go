package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	go func() {
		i := 1
		for {
			i = 1 - i
			if i == 0 {
				s = FIRST
			} else {
				s = SECOND
			}
			time.Sleep(10)
		}
	}()
	for {
		if s == "WHAT" {
			panic(s)
		}
		fmt.Println(s)
		time.Sleep(10)
	}
}

const (
	FIRST = "WHAT THE"
	SECOND = "F*CK"
)
