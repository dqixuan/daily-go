package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {

}

// RPC规则: 方法只能有两个序列化参数，第二个参数是指针， 返回一个error类型参数，方法是公开方法

func (hs *HelloService) Hello(request string, reply *string) (error) {
	*reply = "hello: " + request
	return nil
}

func main() {
	 rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp err:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
