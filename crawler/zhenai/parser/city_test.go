package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseCity(t *testing.T) {
	content, err := http.Get("http://www.zhenai.com/zhenghun/wuxi")
	if err != nil {
		panic("get failed")
	}
	defer content.Body.Close()
	body, err := ioutil.ReadAll(content.Body)
	if err != nil {
		panic("ReadAll failed")
	}
	req := ParseCity(body)
	//fmt.Printf("res=%+v", req)
	for _, re := range req.Requests {
		fmt.Printf("res=%+v\n", re)

	}
}