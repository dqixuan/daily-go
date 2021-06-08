package main

import (
	"context"
	"fmt"
	"go.mod/practice/distributed/registry"
	"log"
	"net/http"
)

/**
  name:丁其轩
  date:2021/6/8
  time:22:55
*/

func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("Registry Service started. Press any key to stop. \n")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<- ctx.Done()
	fmt.Println("Registry service is shutting down.")
}
