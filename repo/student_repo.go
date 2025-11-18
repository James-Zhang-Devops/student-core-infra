package repo

import "student-score/model"

type StudentRepo interface {
	Save(students []model.Student) error
	Load([]model.Student, error)
}