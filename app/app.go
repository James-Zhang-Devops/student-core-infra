package app

import (
	"student-score/config"
	"student-score/repository"
	"student-score/service/impl"
	"student-score/controller"
)

type Application struct {
	Config     *config.AppConfig
	Repo       repository.StudentRepo
	Service    *impl.StudentServiceImpl
	Controller *controller.StudentController
}

func NewApplication(configPath string) (*Application, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	repo := repository.NewFileStudentRepo(cfg.Storage.FilePath)

	studentService := impl.NewStudentService(repo, cfg)

	controller := controller.NewStudentController(studentService)

	return &Application{
		Config:     cfg,
		Repo:       repo,
		Service:    studentService,
		Controller: controller,
	}, nil
}

func (app *Application) Start() error {
	return app.Controller.Run()
}