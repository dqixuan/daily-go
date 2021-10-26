package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(urlString string) (result [] byte, err error) {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		return nil, fmt.Errorf("fetch get failed, err=%w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch ReadAll failed, err=%w", err)
	}
	return
}
