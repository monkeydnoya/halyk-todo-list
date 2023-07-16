package postgres

import (
	"fmt"
	"halyk-todo-list-api/internal/data/database"
	"halyk-todo-list-api/internal/data/database/postgres/entities"
	"halyk-todo-list-api/internal/domain/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func TaskEntityToModel(t entities.Task) models.Task {
	checkList := make([]models.CheckList, 0)
	for _, listItem := range t.CheckList {
		checkList = append(checkList, CheckListEntityToModel(listItem))
	}

	tags := make([]models.Tag, 0)
	for _, tag := range t.Tags {
		tags = append(tags, TagEntityToModel(tag))
	}

	return models.Task{
		Id:          t.Id.String(),
		Name:        t.Name,
		Description: t.Description,
		Deadline:    t.Deadline,
		CreatedAt:   t.CreatedAt,
		FinishedAt:  t.FinishedAt,
		Done:        t.Done,
		CheckList:   checkList,
		Tags:        tags,
	}
}

func TaskModelToEntity(t models.Task) entities.Task {
	id, _ := uuid.FromString(t.Id)
	checkListEntities := make([]entities.CheckList, 0)
	for _, listItem := range t.CheckList {
		checkListEntities = append(checkListEntities, CheckListModelToEntity(listItem))
	}

	tagsEntities := make([]entities.Tag, 0)
	for _, tag := range t.Tags {
		tagsEntities = append(tagsEntities, TagModelToEntity(tag))
	}

	return entities.Task{
		Id:          id,
		Name:        t.Name,
		Description: t.Description,
		Deadline:    t.Deadline,
		CreatedAt:   t.CreatedAt,
		FinishedAt:  t.FinishedAt,
		Done:        t.Done,
		CheckList:   checkListEntities,
		Tags:        tagsEntities,
	}
}

func CheckListModelToEntity(cl models.CheckList) entities.CheckList {
	id, _ := uuid.FromString(cl.Id)
	return entities.CheckList{
		Id:          id,
		Name:        cl.Name,
		Description: cl.Description,
		FinishedAt:  cl.FinishedAt,
		Done:        cl.Done,
	}
}

func CheckListEntityToModel(cl entities.CheckList) models.CheckList {
	return models.CheckList{
		Id:          cl.Id.String(),
		Name:        cl.Name,
		Description: cl.Description,
		FinishedAt:  cl.FinishedAt,
		Done:        cl.Done,
	}
}

func TagModelToEntity(t models.Tag) entities.Tag {
	id, _ := uuid.FromString(t.Id)
	return entities.Tag{
		Id:          id,
		Name:        t.Name,
		Description: t.Description,
	}
}

func TagEntityToModel(t entities.Tag) models.Tag {
	return models.Tag{
		Id:          t.Id.String(),
		Name:        t.Name,
		Description: t.Description,
	}
}

func (*pgConnection) convertError(err error) error {
	switch err {
	case nil:
		return nil
	case gorm.ErrDuplicatedKey:
		return database.ErrDuplicatedKey
	case gorm.ErrPrimaryKeyRequired:
		return database.ErrPrimaryKey
	case gorm.ErrInvalidField:
		return database.ErrInvalidField
	case gorm.ErrInvalidData:
		return database.ErrInvalidData
	case gorm.ErrInvalidTransaction:
		return database.ErrInvalidTx
	}

	err = fmt.Errorf("%s: %s", database.ErrUnknownError, err)
	return err
}
