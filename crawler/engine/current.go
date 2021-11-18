package engine

import (
	"fmt"
	"time"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}


type Scheduler interface {
	Submit(Request)
	ConfigChan(chan Request)
}

func (ce ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	ce.Scheduler.ConfigChan(in)
	for i:=0; i<ce.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	for {
		select {
		case result := <- out:
			for _, item := range result.Item {
				fmt.Printf("Got item: %v\n", item)
			}
			for _, request := range result.Requests {
				ce.Scheduler.Submit(request)
			}
		case <- time.After(30*time.Second):
			fmt.Println("break")
			break
		}


	}

}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <- in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
