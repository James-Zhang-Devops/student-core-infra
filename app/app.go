package append

import (
	"student-score/config"
	"student-score/repo"
	"student-score/service/impl"
	"student-score/controller"
)

type Application struct {
	Config     *config.AppConfig
	Repo       repo.StudentRepo
	Service    *impl.StudentServiceImpl
	Controller *controller.StudentController
}

func NewApplication(config string) (*Application, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	var repo repo.StudentRepo
	switch cfg.Storage.Type {
	case "file":
		repo = repo.NewFileStudentRepo(cfg.Storage.FilePath)
	}

	service := impl.NewStudentService(repo, cfg)

	controller := controller.NewStudentController(service)

	return &Application{
		Config:     cfg,
		Repo:       repo,
		Service:    service,
		Controller: controller,
	}, nil
}

func (app *Application) start() error {
	return app.Controller.Run()
}