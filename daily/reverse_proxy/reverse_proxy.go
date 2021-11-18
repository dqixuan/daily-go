package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

// 获取主机host, 创建一个反向代理
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url2.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errHandler()

	originalDirector := proxy.Director

	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}
	return proxy, nil
}

func originalDirector(req *http.Request) {

}

func modifyRequest(req *http.Request){
	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
}

func modifyResponse() func(*http.Response) error {
	return func(response *http.Response) error {
		response.Header.Set("X-Proxy", "Magical")
		return nil
	}
}

func errHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

// 使用proxy处理请求
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	proxy, err := NewProxy("http://my-api-server.com")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
