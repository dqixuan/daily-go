package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	aChan := make(chan bool, 1)
	bChan := make(chan bool)
	cChan := make(chan bool)

	aChan <- true
	wg.Add(3)
	go printA(wg, aChan, bChan)
	go printB(wg, bChan, cChan)
	go printC(wg, cChan, aChan)

	wg.Wait()
	close(aChan)
	close(bChan)
	close(cChan)
}

func printA(wg *sync.WaitGroup, aChan, bChan chan bool) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<- aChan
		fmt.Println("A")
		bChan <- true
	}
}

func printB(wg *sync.WaitGroup, bChan, cChan chan bool) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<- bChan
		fmt.Println("B")
		cChan <- true
	}
}

func printC(wg *sync.WaitGroup, cChan, aChan chan bool) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<- cChan
		fmt.Println("C")
		aChan <- true
	}
}
