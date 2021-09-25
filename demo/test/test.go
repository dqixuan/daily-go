package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var arr [10]int
	for i:=0; i < 10; i++ {
		go func() {
			for {
				arr[i]++
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Microsecond)
	fmt.Println(arr)
}


