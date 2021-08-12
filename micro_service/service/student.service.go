package service

import (
	"context"
	"errors"
	"go.mod/micro_service/api/student"
	"time"
)

type StudentService struct {
}


func (ss *StudentService) GetStudent(ctx context.Context, req *student.GetStudentReq, resp *student.GetStudentReply) error {
	time.Sleep(30*time.Second)
	studentMap := map[string]student.StudentInfo{
		"davie":  student.StudentInfo{Name: "davie", Grade: 80},
		"steven": student.StudentInfo{Name: "steven",Grade: 90},
		"tony":   student.StudentInfo{Name: "tony", Grade: 85},
		"jack":   student.StudentInfo{Name: "jack", Grade: 96},
	}

	if req.Name == "" {
		return errors.New(" 请求参数错误,请重新请求。")
	}

	s := studentMap[req.Name]

	if s.Name != "" {
		resp.Student = &s
		return nil
	}
	return errors.New(" 未查询当相关学生信息 ")
}
