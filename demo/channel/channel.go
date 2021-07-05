package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 字符串数组， 获取字母个数

func getResult(stringSlice []string) map[int32]int32 {
	result := map[int32]int32{}
	for _, str := range stringSlice {
		for _, r := range str {
			result[r]++
		}
	}
	return result
}

func main1() {
	s := paramHelper()
	now := time.Now()
	// 单协程执行
	fmt.Println(getResult(s))
	fmt.Println(time.Since(now))

	// 多协程执行
	now = time.Now()
	ch := make(chan map[int32]int32, 10)
	result := make(map[int32]int32)
	go concurrentGetResult(s, 10, ch)
	for {
		select {
		case m:= <- ch:
			for key, val := range m {
				result[key] += val
			}
		case <-time.After(10*time.Millisecond):
			fmt.Println(result)
			fmt.Println(time.Since(now))
			return
		}
	}

}

func concurrentGetResult(stringSlice []string, concurrency int, ch chan  map[int32]int32) {
	product := make(chan string, 10)

	var wg sync.WaitGroup
	for i:=0 ; i < concurrency; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, no int) {
			defer wg.Done()
			for c := range product {
				//fmt.Printf("%d goroutine worker %s\n", no, c)
				m := helper(c)
				ch <- m
			}
		}(&wg, i)
	}


	for _, str := range stringSlice {
		product <- str
	}

	wg.Wait()
	close(product) // product中的数据消费完， 关闭channel
	close(ch) // 写入数据方，关闭channel
}

func helper(str string) map[int32]int32 {
	result := map[int32]int32{}
	for _, r := range str {
		result[r]++
	}
	return result
}



func paramHelper() []string {
	st := []string{}
	for i:=0; i < 2000000; i++ {
		str := ""
		for j :=0; j < 20; j++ {
			str += strconv.Itoa(i)
		}
		st = append(st, str)
	}
	return st
}