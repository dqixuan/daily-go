package main

import (
	origin_plus_grpc_client "daily-go/grpc/origin_plus_grpc/client"
	"fmt"
	"log"
)

func main() {
	client, err := origin_plus_grpc_client.DialHelloService("tcp", "localhost:1223")
	if err != nil {
		log.Fatal("DialHelloService failed, err:", err)
	}
	var reply string
	err = client.Hello("test hello", &reply)
	if err != nil {
		log.Fatal("client Hello error:", err)
	}
	fmt.Println(reply)
}
