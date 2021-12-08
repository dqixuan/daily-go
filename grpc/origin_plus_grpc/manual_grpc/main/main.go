package main

import (
	"daily-go/grpc/origin_plus_grpc/manual_grpc"
	"daily-go/grpc/origin_plus_grpc/manual_grpc/hello_service"
	"log"
	"net"
	"net/rpc"
)



func main() {
	manual_grpc.RegisterHelloService(new(hello_service.HelloService))
	listener, err := net.Listen("tcp", ":1223")
	if err != nil {
		log.Fatal("net Listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listerner Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
