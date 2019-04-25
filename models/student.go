package models

import (
	"errors"
)

var (
	StudentList map[string]*Student
	SerialNo    int64
)

type Student struct {
	StudentId   string
	StudentName string
	ClassName   string
}

func init() {
	StudentList = make(map[string]*Student)
	StudentList["s201700001"] = &Student{"s201700001", "Sharon", "apple"}
	StudentList["s201700002"] = &Student{"s201700002", "Summer", "apple"}

	SerialNo = 201700018
}

func GetStudent(studentId string) (student *Student, err error) {
	if v, ok := StudentList[studentId]; ok {
		return v, nil
	}
	return nil, errors.New("studentId Not Exist")
}

func GetAllStudents() map[string]*Student {
	return StudentList
}
