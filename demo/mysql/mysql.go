package main

import "github.com/micro/go-micro/util/log"

func main() {
	InitDb()
	result, err := getUsrs(DB)
	if err != nil {
		log.Errorf("getUsers failed, err=%v", err)
		return
	}
	for i, val := range result {
		log.Infof("user%d = %+v", i, val)
	}
}

