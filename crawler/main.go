package main

import (
	"go.mod/crawler/engine"
	"go.mod/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}


