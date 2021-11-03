package scheduler

import (
	"go.mod/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler) Submit(request engine.Request) {
	s.workerChan <- request
}

func (s *SimpleScheduler) ConfigChan(c chan engine.Request) {
	s.workerChan = c
}


