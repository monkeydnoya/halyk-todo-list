package todolist

import (
	"halyk-todo-list-api/internal/config"
	"halyk-todo-list-api/internal/data/database"
	"halyk-todo-list-api/internal/service"

	"go.uber.org/zap"
)

type Service struct {
	log    *zap.SugaredLogger
	Config config.Config

	Db database.Database
}

func NewService(config config.Config) service.TodoListService {
	var database database.Database
	var err error

	ServiceContext := Service{
		Config: config,
		log:    zap.NewExample().Sugar(),
	}
	switch config.Database.DatabaseType {
	case "postgres":
		database, err = config.Database.Postgres.NewConnection()
		if err != nil {
			ServiceContext.log.Errorf("postgres database: %v", err)
			panic(err)
		}
		ServiceContext.log.Infof("postgres database: connected successfully: %s:%s", config.Database.Postgres.Host, config.Database.Postgres.Port)
	default:
		panic("unknown database type: " + config.Database.DatabaseType)
	}
	if err := database.Start(); err != nil {
		ServiceContext.log.Errorf("postgres database: could not start connection: %s", err)
		panic(err)
	}
	ServiceContext.Db = database
	return &ServiceContext
}
