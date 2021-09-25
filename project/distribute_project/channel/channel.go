package main

import (
	"fmt"
	"time"
)

func main() {
	//channelDemo()
	//chanDemo()

	//bufferedChannel()
	channelClose()
	time.Sleep(time.Microsecond)
}

func createWorker1(id int ) chan <- int {
	c := make(chan int)
	go func() {
		for {
			n := <- c
			fmt.Printf("Worker %d received %c\n", id, n)
		}
	}()
	return c
}

func channelDemo() {
	var channels [10]chan <-int
	for i:=0; i<10; i++ {
		channels[i] = createWorker1(i)
	}
	for i:=0; i< 10; i++ {
		channels[i] <- 'a' + i
	}
	for i:=0; i< 10; i++ {
		channels[i] <- 'A' + i
	}
}

func worker(id int, c chan int) {
	for n := range c{
		fmt.Printf("worker %d received %c\n", id, n)
	}
}
func workerOk(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			return
		}
		fmt.Printf("worker %d received %c\n", id, n)
	}
}

func chanDemo() {
	var channels [10] chan int
	for i:=0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

// close(channel)  发送方关闭channel


func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Second)
}

func channelClose() {
	c := make(chan int, 3)
	go workerOk(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Second)
}
