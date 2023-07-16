package postgres

import (
	"fmt"
	"halyk-todo-list-api/internal/data/database/postgres/entities"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (pg *pgConnection) Start() error {
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		url.QueryEscape(pg.conf.Username),
		url.QueryEscape(pg.conf.Password),
		url.QueryEscape(pg.conf.Host),
		pg.conf.Port,
		url.QueryEscape(pg.conf.DbName),
	)
	pg.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		pg.log.Errorw("postgres: failed to create new posgresql connection: %v", err)
		return err
	}

	if err = pg.db.AutoMigrate(&entities.Task{}); err != nil {
		pg.log.Errorw("postgres: failed to migrate task entity: %v", err)
		return err
	}

	if err = pg.db.AutoMigrate(&entities.CheckList{}); err != nil {
		pg.log.Errorw("postgres: failed to migrate check list entity: %v", err)
		return err
	}

	if err = pg.db.AutoMigrate(&entities.Tag{}); err != nil {
		pg.log.Errorw("postgres: failed to migrate tag entity: %v", err)
		return err
	}
	return nil
}

func (pg *pgConnection) Stop() error {
	dbInstance, err := pg.db.DB()
	if err == nil {
		err = dbInstance.Close()
		if err != nil {
			pg.log.Errorw("postgres: close connection error: %v", err)
		}
	}
	pg.log.Info("postgres: connection closed")
	return nil
}
