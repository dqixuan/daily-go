package registry

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

/**
  name:丁其轩
  date:2021/6/8
  time:22:41
*/


const (
	ServerPort = ":4000"
	ServiceUrl = "http://localhost"+ServerPort + "/services"
)

type registry struct {
	registration []Registration
	mutex *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registration = append(r.registration, reg)
	r.mutex.Unlock()
	return nil
}

func (r *registry) remove(serviceUrl string) error {
	for i := range r.registration {
		if r.registration[i].ServiceUrl == serviceUrl {
			r.mutex.Lock()
			r.registration = append(r.registration[:i], r.registration[i+1:]...)
		}
	}
	return fmt.Errorf("Service at URL %s not found", serviceUrl)
}

var reg = registry{
	registration: make([]Registration, 0),
	mutex:   	  new(sync.Mutex),
}

type RegistryService struct {
}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received.")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %s\n", r.ServiceName, r.ServiceUrl)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}