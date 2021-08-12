package main

import (
	"github.com/micro/go-micro"
	"go.mod/micro_service/api/student"
	"go.mod/micro_service/service"
	"log"
	"time"
)

func main() {
	sv := micro.NewService(
		micro.Name("student"),
		micro.RegisterTTL(10*time.Second),
		micro.RegisterInterval(5*time.Second),
		)

	sv.Init()

	student.RegisterStudentHandler(sv.Server(), new(service.StudentService))

	err := sv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
