package impl

import (
	"student-score/model"
	"student-score/repository"
	"student-score/config"
)

type StudentServiceImpl struct {
	repo repository.StudentRepo
	config *config.AppConfig
}

func NewStudentService(repo repository.StudentRepo, cfg *config.AppConfig) *StudentServiceImpl {
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
		if score >= level.Min && score <= level.Max {
			return level.Name
		}
	}

	return "Unknow Grade"
}

func (s *StudentServiceImpl) ListStudent() ([]model.Student, error) {
    return s.repo.Load()
}

func (s *StudentServiceImpl) GetStudent(id int) (*model.Student, error) {
    students, err := s.repo.Load( )
	if err != nil {
		return nil, err
	}

	for _, student := range students {
		if student.ID == id {
			return &student, nil
		}
	}
    return nil, nil
}