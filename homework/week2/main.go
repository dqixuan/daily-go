package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		http.HandleFunc("/hello", hello)
		return http.ListenAndServe(":8888", nil)
	})


	g.Go(func() error {
		signalChan := make(chan os.Signal)

		go func() {
			signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
			fmt.Println("signal received")
		}()
		<-signalChan
		return errors.New("test")
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello"))
}
