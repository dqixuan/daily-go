package main

import (
	"fmt"
	"net"

	"go.mod/demo/tcp_stick_package/my_protocol"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8089")
	if err != nil {
		fmt.Println("net Dial failed, err:", err)
		return
	}
	protocol := my_protocol.MyProtocol{}
	for i := 0; i < 10; i++ {
		dataBytes, err := protocol.Encode("solve tcp stick package")
		if err != nil {
			fmt.Printf("encode data failed, err=%v\n", err)
			return
		}
		_, err = conn.Write(dataBytes)
		if err != nil {
			fmt.Printf("conn write failed, err=%v\n", err)
			return
		}
	}
}
