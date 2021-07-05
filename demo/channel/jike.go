package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
  name:丁其轩
  date:2021/6/29
  time:21:19
*/

/**
	1. 什么时候停止
	2. 怎么停止
 */





func main3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello GopherCon SG")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
/**
func leak() {
	go func() {
		record, err := search("term")
		ch <- result{record, err}
	}()
	select {
	case <- ctx.Done:

	case result := <- ch:
		return
	}
}
 */

/**
	Leave concurrency to the caller
	1 func ListDirectory(dir string) ([]string, error)
		同步接口， 目录太大，遍历花费大量时间
	2 func ListDirectory(dir string) chan string
		函数内部启动goroutine， 关闭通道(结束、报错)，接受者必须读完整个目录，无法中断。
	3 func ListDirectory(dir string, fn func(string))
		filepath.WalkDir,  将异步执行函数的决定权交给业务调用方
 */

func main2() {
	// keep yourself busy or do the work yourself  脏活累活自己做
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello GopherCon SG")
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	// 防止主线程退出
	select {
	}
}
