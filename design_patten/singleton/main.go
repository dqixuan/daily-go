package main

import (
	"fmt"
	"sync"
)
/*
	sync.Once结构体中有：done标志， mutex锁
	Do方法：1、先原子性判断done是否为0，
		=0， 已执行过Do方法，直接退出
		!=0，执行doSlow方法
	func (o *Once) doSlow(f func()) {
		加锁，是为了确保多个请求同时执行Do方法，只执行一次f()
		o.m.Lock()
		defer o.m.Unlock()
		if o.done == 0 {
			defer atomic.StoreUint32(&o.done, 1)
			f()
		}
	}
*/

func main() {
	obj1 := GetInstance()
	fmt.Printf("pointer=%p\n", obj1)
	obj2 := GetInstance()
	fmt.Printf("pointer=%p\n", obj2)
	obj3 := GetInstance()
	fmt.Printf("pointer=%p\n", obj3)

	/*
		pointer=0x11a7c70
		pointer=0x11a7c70
		pointer=0x11a7c70
	*/
}

var (
	once sync.Once
)

// Singleton 单例模式
type Singleton struct {
}

var s *Singleton

// GetInstance 获取实例
func GetInstance() *Singleton {
	if s == nil {
		once.Do(
			func() {
			s = &Singleton{}
		})
	}
	return s
}
