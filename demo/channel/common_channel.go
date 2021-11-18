package main

import "fmt"

/**
	nil channel  	读写一直阻塞
	带缓存的channel	读:读取对应类型零值，写: panic  send on closed channel
 */


func main() {
	c := make(chan int, 1)
	close(c)
	c<-1
	fmt.Println(<-c)
}
