package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var timeout = time.Second * 3

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{"172.18.3.11:6379", "172.18.3.11:6381", "172.18.3.11:6380"},
	})
	err := client.Ping().Err()
	if err == nil {
		fmt.Println("Redis cluster is ok")
	} else {
		fmt.Println("Redis cluster is not ok. err=%v", err)
	}
}
