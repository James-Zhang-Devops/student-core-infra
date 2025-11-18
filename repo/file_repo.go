package repo

import (
	"encoding/json"
	"os"
	"student-score/model"
)

type FileStudentRepo struct {
	filename string
}

func NewFileStudentRepo(filename string) *FileStudentRepo {
	return &FileStudentRepo{filename: filename}
}

func (r *FileStudentRepo) Save(students []model.Student) error {
	data, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filename, data, 0644)
}

func (r *FileStudentRepo) Load() ([]model.Student, error) {
	data, err := os.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Student{}, nil
		}
		return nil, err
	}

	var students []model.Student
	if err := json.Unmarshal(data, &students); err != nil {
		return nil, err
	}
	return students, nil
}