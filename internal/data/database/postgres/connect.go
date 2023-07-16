package postgres

import (
	"halyk-todo-list-api/internal/data/database"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type pgConnection struct {
	conf PostgresConfig
	log  *zap.SugaredLogger
	db   *gorm.DB
}

func (pConf *PostgresConfig) NewConnection() (database.Database, error) {
	pg := pgConnection{
		conf: *pConf,
		log:  zap.NewExample().Sugar(),
	}

	return &pg, nil
}
