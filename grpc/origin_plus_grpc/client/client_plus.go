package origin_plus_grpc_client

import (
	"daily-go/grpc/origin_plus_grpc/manual_grpc"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

// DialHelloService 实现rpc
func DialHelloService(network, addr string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(manual_grpc.HelloServiceName+".Hello", request, &reply)
}

// HelloServiceClient  结构体实现了HelloServiceInterface接口
var _ manual_grpc.HelloServiceInterface = (*HelloServiceClient)(nil)
