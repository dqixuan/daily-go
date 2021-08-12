package main

import (
	"github.com/micro/go-micro"
	"go.mod/micro_service/api/student"
	"go.mod/micro_service/service"
)

func main() {
	sevice := micro.NewService(
		micro.Name("student"))

	sevice.Init()

	student.RegisterStudentHandler(sevice.Server(), new(service.StudentService))
}
