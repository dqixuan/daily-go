package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main1() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func HelloServer(r http.ResponseWriter, req *http.Request) {
	_, _ = r.Write([]byte("hello world\n"))
}

func main() {
	//resp, err := http.Get("http://imooc.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//s, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(s))

	request, err := http.NewRequest(http.MethodGet, "http://imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent","Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Mobile Safari/537.36")
	client1 := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	response, err := client1.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	s, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
