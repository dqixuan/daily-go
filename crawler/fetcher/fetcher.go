package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var ticker = time.NewTicker(time.Microsecond*100)

func Fetch(urlString string) (result [] byte, err error) {
	<- ticker.C
	resp, err := http.Get(urlString)
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
