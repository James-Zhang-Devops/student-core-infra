package service

import "student-score/model"

type StudentService interface {
	AddStudent(student model.Student) error
	GetStudent(id int) (*model.Student, error)
	ListStudent() ([]model.Student, error)
	GetGradeLevel(score float64) string
}