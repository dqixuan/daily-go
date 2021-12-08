package main

import (
	"fmt"
	consular "github.com/hashicorp/consul/api"
	"log"
	"net"
	"time"
)

func main() {
	client, err := consular.NewClient(consular.DefaultConfig())

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	for {

		time.Sleep(time.Second * 3)
		var services map[string]*consular.AgentService
		var err error

		services, err = client.Agent().Services()

		if nil != err {
			log.Println("in consual list Services:", err)
			continue
		}

		if _, found := services["node01"]; !found {
			log.Println("node01 not found")
			continue
		}

		sendData(services["node01"])
	}
}

func sendData(service *consular.AgentService) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", service.Address, service.Port))

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	i := 0
	for {
		i++
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}

		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		log.Println("get:", string(buf[0:n]))

		//等一秒钟
		time.Sleep(time.Second)
	}
}