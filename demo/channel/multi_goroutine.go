package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

/**
  name:丁其轩
  date:2021/6/29
  time:21:43
*/


// 1、将并发的选择性， 交给业务调用方

func serveApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello QCon!")
	})
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func serveDebug() {
	http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func mainA() {
	go serveDebug()  // serveDebug 意外退出 主线程不知道
	serveApp()       // serveApp 退出，程序结束
}

// 2、希望 serveDebug 和 serveApp 同时退出。不建议用log.Fatal， 调用os.Exit()。defer函数都不会执行

func serveAppB() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serveDebugB() {
	if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func mainB() {
	go serveDebugB()  // serveDebug 意外退出 主线程不知道
	go serveAppB()       // serveApp 退出，程序结束
	select {}
}

//
func serve(addr string, handler http.Handler, stop <- chan struct{}) error {
	s := http.Server{
		Addr:              addr,
		Handler:           handler,
	}
	go func() {
		<- stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func mainC() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	//go func() {
	//	done <- serveDebugC(stop)
	//}()
	//go func() {
	//	done <- serveAPPC(stop)
	//}()

	var stopped bool
	for i:=0; i < cap(done); i++ {
		if err := <- done; err != nil {
			fmt.Printf("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
