package config

import (
	"halyk-todo-list-api/internal/config"
)

func Loader(cfg *config.Config) {
	cfg.Database.DatabaseType = Config("DATABASE_TYPE")
	if cfg.Database.DatabaseType != "" {
		switch cfg.Database.DatabaseType {
		case "postgres":
			cfg.Database.Postgres.DbName = Config("POSTGRES_DATABASE_DBNAME")
			cfg.Database.Postgres.Host = Config("POSTGRES_DATABASE_HOST")
			cfg.Database.Postgres.Port = Config("POSTGRES_DATABASE_PORT")
			cfg.Database.Postgres.Username = Config("POSTGRES_DATABASE_USERNAME")
			cfg.Database.Postgres.Password = Config("POSTGRES_DATABASE_PASSWORD")
		}
	}

	cfg.Api.Name = Config("API_NAME")
	cfg.Api.Host = Config("API_HOST")
	cfg.Api.Port = Config("API_PORT")
}
