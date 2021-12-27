package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
	数据量    延迟    	加锁成功    失败
	10		 34.591µs	10		   0
	100		 179.452µs  96		   4
	1000	 1.201808ms 741		   243
	10000	 4.924594ms 5694	   3789
 */

type MyLock struct {
	c chan struct{}
}

// 利用channel控制并发
func NewLock() MyLock {
	var l = MyLock{}
	l.c = make(chan struct{}, 1)
	l.c <- struct {}{}
	return l
}

// 从channel中读取数据视为加锁成功
func (ml MyLock) Lock() bool {
	lockResult := false
	select {
	case <- ml.c:
		lockResult = true
	default:

	}
	return lockResult
}

// 向channel中发送数据视为解锁
func (ml MyLock) Unlock() {
	ml.c <- struct{}{}
}


func main() {
	counter := 0

	success, failure := 0, 0
	l := NewLock()
	now := time.Now()
	var sw sync.WaitGroup
	for i:=0; i < 10000; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			if !l.Lock() {
				failure++
				return
			}
			counter++
			success++
			l.Unlock()
		}()
	}
	sw.Wait()
	fmt.Println("latency time:", time.Since(now))
	fmt.Println("success:", success)
	fmt.Println("failure:", failure)
	fmt.Println(runtime.NumGoroutine())
}


