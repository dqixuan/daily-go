package engine

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"go.mod/crawler/fetcher"
)

type SimpleEngine struct {
}

func (se SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) != 0 {
		r := requests[0]
		log.Infof("fetching url:%s", r.Url)
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Item {
			log.Infof("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	text, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Errorf("fetcher url failed, err=%v", err)
		return ParseResult{}, err
	}
	fmt.Println(len(text))
	parseResult := r.ParserFunc(text)
	return parseResult, nil
}
