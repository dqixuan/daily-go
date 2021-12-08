package main

import (
	"log"
	"net"
	"net/rpc"
)

// rpc接口规范：两个序列化参数，第二个参数是指针类型；返回值只有一个error; 方法必须是公开的。

type HelloService struct {
}

func (hs *HelloService) Hello(request string, reply  *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1222")
	if err != nil {
		log.Fatal("net Dial error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("listen Accept error:", err)
	}
	rpc.ServeConn(conn)
}
