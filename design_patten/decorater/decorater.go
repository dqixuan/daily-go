package main

import (
	"fmt"
	"net/http"
)

/*
装饰器模式：包装现有方法，允许在现有方法的执行之前或之后添加自定义功能，以达到功能增强的目的。
*/

func logger(m func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logger start")
		m(w, r)
		fmt.Println("logger end")
	})
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", logger(handlerFunc))
	http.ListenAndServe(":8080", nil)
}
