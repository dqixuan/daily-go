package parser

import (
	"fmt"
	"go.mod/crawler/engine"
	"regexp"
)

//
//func ParseProfile(content []byte) engine.ParseResult {
//	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
//	if err != nil {
//		panic(fmt.Sprintf("ParseProfile  NewDocumentFromReader err=%v", err))
//	}
//	fmt.Println("hldjsljf")
//	dom.Find("div[class=m-btn purple]").Each(func(i int, selection *goquery.Selection) {
//		fmt.Println(selection.Text())
//	})
//	fmt.Println("hldjsljf")
//	return engine.ParseResult{}
//}

var (
	personReg = regexp.MustCompile(`<div class="m-btn purple"[^>]+>([^<]+)</div>`)
)

func ParseProfile(content []byte) engine.ParseResult {
	all := personReg.FindAllSubmatch(content, -1)
	fmt.Println("all=", len(all))
	for _ , val := range all {
		fmt.Printf("val=%s", string(val[1]))
	}
	return engine.ParseResult{}
}