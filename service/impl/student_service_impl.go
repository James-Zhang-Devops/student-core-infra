package impl

import (
	"student-score/model"
	"student-score/repo"
	"student-score/config"
)

type StudentServiceImpl struct {
	repo repo.StudentRepo
	config *config.AppConfig
}

func NewStudentService(repo repo.StudentRepo, cfg *config.AppConfig) *StudentServiceImpl {
	return &StudentServiceImpl {
		repo:   repo,
		config: cfg,
	}
}

func (s *StudentServiceImpl) AddStudent(student model.Student) error {
	students, err := s.repo.Load()
	if err != nil {
		return err
	}

	if student.ID == 0 {
		student.ID = len(students) + 1
	}

	students = append(students, student)
	return s.repo.Save(students)
}

func (s *StudentServiceImpl) GetGradeLevel(score float64) string {
	for _, level := range s.config.GradeConfig.Levels {
		if score >= level.Min && socre <= level.Max {
			return level.Name
		}
	}

	return "Unknow Grade"
}