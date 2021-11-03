package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseProfile(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://album.zhenai.com/u/1526073307", nil)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36")

	content, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("get failed")
	}
	defer content.Body.Close()
	body, err := ioutil.ReadAll(content.Body)
	if err != nil {
		panic("ReadAll failed")
	}
	fmt.Println(string(body))
	fmt.Println(len(body))
	req1 := ParseProfile(body)
	//fmt.Printf("res=%+v", req)
	for _, re := range req1.Requests {
		fmt.Printf("res=%+v\n", re)
	}
}
