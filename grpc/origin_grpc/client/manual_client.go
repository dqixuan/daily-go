package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1222")
	if err != nil {
		log.Fatal("client Dial error:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello",&reply)
	if err != nil {
		log.Fatal("client Call error:", err)
	}
	fmt.Println(reply)
}
