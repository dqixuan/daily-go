package main

import (
	"encoding/json"
	"fmt"
)

/**
1、只有导出的字段才能被序列化(字段名称首字母大写)
*/

func main() {
	u := person{
		Name:  "ding",
		Age:   10,
		array: []string{"1", "2", "4"},
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}

type person struct {
	Name  string `json:"name"`
	Age   int
	array []string
}
