package parser

import (
	"go.mod/crawler/engine"
	"regexp"
)

const urlRegexp = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(content []byte) (result engine.ParseResult) {
	reg := regexp.MustCompile(urlRegexp)
	result = engine.ParseResult{}

	all := reg.FindAllSubmatch(content, -1)
	for _, v := range all {
		result.Item = append(result.Item, string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: engine.NilRequest,
		})
	}
	return
}
