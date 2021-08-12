package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"go.mod/micro_service/api/student"
	"log"
)

func main() {
	sv := micro.NewService(
		micro.Name("student.client"),
		)

	sv.Init()

	studentService := student.NewStudentClient("student", sv.Client())

	resp, err := studentService.GetStudent(context.TODO(), &student.GetStudentReq{Name:"davie"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp=%+v", resp)
}
