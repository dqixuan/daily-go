package parser

import (
	"fmt"
	"go.mod/crawler/engine"
	"regexp"
)

const people = `<a href="(http://album.zhenai.com/u/[0-9]{10})"[^>]+>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	reg := regexp.MustCompile(people)
	matches := reg.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, match := range matches {
		fmt.Printf("match=%s\n", match[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: engine.NilRequest,
		})
		result.Item = append(result.Item, string(match[2]))
	}
	return result
}
