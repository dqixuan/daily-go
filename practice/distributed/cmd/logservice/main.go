package main

import (
	"context"
	"fmt"
	"go.mod/practice/distributed/log"
	"go.mod/practice/distributed/registry"
	"go.mod/practice/distributed/service"
	stdlog "log"
)

/**
  name:丁其轩
  date:2021/6/8
  time:22:17
*/

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "3000"
	serviceUrl := fmt.Sprintf("http://%s:%s", host, port)
	reg := registry.Registration{
		ServiceName: "Log Service",
		ServiceUrl:  serviceUrl,
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		reg,
		log.RegisterHandlers,
		)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<- ctx.Done()
	fmt.Println("Shutting down log service.")
}
