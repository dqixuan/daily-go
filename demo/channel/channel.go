package main

import (
	"fmt"
	"time"
)

func worker(taskCh <- chan int) {
	const N = 5 
	for i:=0; i < N; i++ {
		go func(id int) {
			for {
				task := <- taskCh
				fmt.Println("finish task: %d by worker %d", task, i)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func main() {
	tashCh := make(chan int, 100)
	go worker(tashCh)
	
	for i:=0; i < 20; i++ {
		tashCh <- i
	}
	select {
	case <- time.After(time.Minute):
		
	}
}
