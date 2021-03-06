package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"go.mod/demo/tcp_stick_package/my_protocol"
)

func main() {
	lsn, err := net.Listen("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println("net listen failed, err:", err)
		return
	}
	defer lsn.Close()

	for {
		conn, err := lsn.Accept()
		if err != nil {
			fmt.Println("conn Accept failed, err:", err)
			return
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	protocol := my_protocol.MyProtocol{}
	for {
		msg, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("process decode failed, err:", err)
			return
		}
		fmt.Printf("data = %s\n", msg)
	}
}
