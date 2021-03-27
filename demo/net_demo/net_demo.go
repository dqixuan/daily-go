package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, today is ", time.Now())
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
}
