package postgres

import (
	"context"
	"halyk-todo-list-api/internal/data/database/postgres/entities"
	"halyk-todo-list-api/internal/domain/models"
)

func (pg *pgConnection) GetTag(ctx context.Context, id string) (models.Tag, error) {
	var tagEntity entities.Tag

	query := pg.db.WithContext(ctx).Model(&entities.Tag{})
	if err := query.Where("id = ?", id).First(&tagEntity).Error; err != nil {
		pg.log.Errorw("postgres: could not found", err)
		return models.Tag{}, pg.convertError(err)
	}

	return TagEntityToModel(tagEntity), nil
}

func (pg *pgConnection) GetTags(ctx context.Context) ([]models.Tag, error) {
	var tagEntities []entities.Tag

	query := pg.db.WithContext(ctx).Model(&entities.Tag{})
	if err := query.Find(&tagEntities).Error; err != nil {
		pg.log.Errorw("postgres: could not found", err)
		return []models.Tag{}, pg.convertError(err)
	}

	var tags []models.Tag
	for _, tagEntity := range tagEntities {
		tags = append(tags, TagEntityToModel(tagEntity))
	}
	return tags, nil
}

func (pg *pgConnection) CreateTag(ctx context.Context, tag models.CreateTag) (models.Tag, error) {
	tagEntity := entities.Tag{
		Name:        tag.Name,
		Description: tag.Description,
	}

	tx := pg.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to create tag")
		}
	}()
	if err := tx.Create(&tagEntity).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to create tag: %v", err)
		return models.Tag{}, pg.convertError(err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to create tag: %v", err)
		return models.Tag{}, pg.convertError(err)
	}

	return TagEntityToModel(tagEntity), nil
}

func (pg *pgConnection) UpdateTag(ctx context.Context, id string, tag models.UpdateTag) error {
	tx := pg.db.WithContext(ctx).Begin()
	query := tx.Model(&entities.Tag{}).Where("id = ?", id)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to create tag")
		}
	}()

	if tag.Name != "" {
		if err := query.Update("name", tag.Name).Error; err != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: failed to update tag: %v", err)
			return pg.convertError(err)
		}
	}
	if tag.Description != "" {
		if err := query.Update("description", tag.Description).Error; err != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: failed to update tag: %v", err)
			return pg.convertError(err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to update tag: %v", err)
		return pg.convertError(err)
	}

	return nil
}

func (pg *pgConnection) DeleteTag(ctx context.Context, tag models.Tag) error {
	tagEntity := TagModelToEntity(tag)

	tx := pg.db.WithContext(ctx).Begin()
	query := tx.Model(&entities.Tag{}).Where("id = ?", tag.Id)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			pg.log.Errorw("postgres: recover transaction: failed to delete tag")
		}
	}()
	if err := query.Delete(&tagEntity).Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete tag: %v", err)
		return pg.convertError(err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		pg.log.Errorw("postgres: failed to delete tag: %v", err)
		return pg.convertError(err)
	}

	return nil
}
