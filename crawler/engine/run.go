package engine

import (
	"github.com/micro/go-micro/util/log"
	"go.mod/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) != 0 {
		r := requests[0]
		log.Infof("fetching url:%s", r.Url)
		requests = requests[1:]
		text, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Errorf("fetcher url failed, err=%v", err)
			continue
		}
		parseResult := r.ParserFunc(text)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Item {
			log.Infof("Got item %v", item)
		}
	}
}
