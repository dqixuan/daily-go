package main

import (
	"fmt"
	consular "github.com/hashicorp/consul/api"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	// 注册服务
	go registerService()

	// 生成一个服务：服务1
	ln, err := net.Listen("tcp", "0.0.0.0:9527")
	if err != nil {
		panic("error：" + err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("error：" + err.Error())
		}
		go EchoServer(conn)
	}

}

func EchoServer(conn net.Conn) {

	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			log.Println("get and echo:", "EchoServer "+string(buf[0:n]))
			conn.Write(append([]byte("EchoServer "), buf[0:n]...))
		case io.EOF:
			log.Printf("Warning: End of data: %s\n", err)
			return
		default:
			log.Printf("Error: Reading data: %s\n", err)
			return
		}
	}
}

func registerService() {
	config := consular.DefaultConfig()
	client, err := consular.NewClient(config)

	if err != nil {
		log.Fatal("consul client error: ", err)
		return
	}

	checkPort := 8080


	registration := new(consular.AgentServiceRegistration)
	registration.ID = "node01"
	registration.Name = "serverNode"
	registration.Port = 9527
	registration.Tags = []string{"serverNode"}
	registration.Address = "127.0.0.1"

	registration.Check = &consular.AgentServiceCheck{
		HTTP: fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout: "3s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
		return
	}
	http.HandleFunc("/check", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println( "consulCheck: ", w)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
}
