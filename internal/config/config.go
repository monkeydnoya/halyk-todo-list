package config

import "halyk-todo-list-api/internal/data/database/postgres"

type Config struct {
	Api
	Database
}

type Database struct {
	DatabaseType string
	Postgres     postgres.PostgresConfig
}

type Api struct {
	Name string
	Host string
	Port string
}
