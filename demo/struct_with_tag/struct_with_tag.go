package main

import (
	"encoding/json"
	"fmt"
)

/**
1、只有导出的字段才能被序列化(字段名称首字母大写)
2、字段首字母大写， 加tag,序列化显示的是jsontag对应名称；不加tag，显示和字段名称相同的
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
