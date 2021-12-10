package functional_options_pattern

import "time"

/**
	使用场景：
	构造函数或是公共API的可变参数， 用于扩展参数(三个或三个以上的参数)

	策略三种：
	1、定义多个构造函数
	2、定义一个新的config结构体来保存
	3、functional option pattern模式

	参考链接：
 */

type server struct {
	host string
	port int32
	timeout time.Duration
	maxIdle int32
}

func New(host string, port int32) *server {
	return &server{
		host:    host,
		port:    port,
	}
}

func NewWithTimeout(host string, port int32, timeout time.Duration) *server {
	return &server{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

func NewWithTimeoutAndMaxIdle(host string, port, maxIdle int32, timeout time.Duration) *server {
	return &server{
		host:    host,
		port:    port,
		timeout: timeout,
		maxIdle: maxIdle,
	}
}

// 使用专门配置的结构体


// 使用functional option pattern
// 定义Option函数类型
type Option func(*server)

func NewByOption(options ...Option) *server {
	svr := &server{}
	for _, f := range options {
		f(svr)
	}
	return svr
}

func WithHost(host string) Option {
	return func(s *server) {
		s.host = host
	}
}

func WithPort(port int32) Option {
	return func(s *server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.timeout = timeout
	}
}

func WithMaxIdle(maxIdle int32) Option {
	return func(s *server) {
		s.maxIdle = maxIdle
	}
}

