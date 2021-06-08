package registry

/**
  name:丁其轩
  date:2021/6/8
  time:22:37
*/

type Registration struct {
	ServiceName ServiceName
	ServiceUrl string
}

type ServiceName string

const (
	LogService = ServiceName("Log Service")
)
