package config

import "student-score/model"

type AppConfig struct {
	GradeConfig model.GradeConfig `json:"grade_config"`
	Storage		StorageConfig     `json:"storage"`
	UI			UIConfig		  `json:"ui"`
}

type StorageConfig struct {
	Type 	 string `json:"type"`
	FilePath string `json:"file_path"`
}

type UIConfig struct {
	Language string `json:"language"`
	Theme	 string `json:"theme"`
}