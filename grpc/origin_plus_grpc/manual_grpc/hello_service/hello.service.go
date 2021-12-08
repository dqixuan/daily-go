package hello_service


type HelloService struct {
}

func (hs *HelloService) Hello(request string, reply  *string) error {
	*reply = "hello: " + request
	return nil
}
