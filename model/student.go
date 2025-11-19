package model

type Student struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
	Class string  `json:"class"`
}

type GradeConfig struct {
	Levels []GradeLevel `json:"levels"`
}

type GradeLevel struct {
	Name  string  `json:"name"`
	Min	  float64 `json:"min"`
	Max   float64 `json:"max"`
	Color string  `json:"color"`
}