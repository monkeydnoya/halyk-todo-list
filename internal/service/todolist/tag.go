package todolist

import (
	"context"
	"halyk-todo-list-api/internal/domain/models"
)

func (s *Service) GetTag(ctx context.Context, id string) (models.Tag, error) {
	tag, err := s.Db.GetTag(ctx, id)
	if err != nil {
		return models.Tag{}, err
	}

	return tag, nil
}

func (s *Service) GetTags(ctx context.Context) ([]models.Tag, error) {
	tags, err := s.Db.GetTags(ctx)
	if err != nil {
		return []models.Tag{}, err
	}

	return tags, nil
}

func (s *Service) CreateTag(ctx context.Context, tag models.CreateTag) (models.Tag, error) {
	createdTag, err := s.Db.CreateTag(ctx, tag)
	if err != nil {
		return models.Tag{}, err
	}

	return createdTag, nil
}

func (s *Service) UpdateTag(ctx context.Context, id string, tag models.UpdateTag) error {
	err := s.Db.UpdateTag(ctx, id, tag)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTag(ctx context.Context, id string) error {
	tag, err := s.Db.GetTag(ctx, id)
	if err != nil {
		return err
	}

	err = s.Db.DeleteTag(ctx, tag)
	if err != nil {
		return err
	}

	return nil
}
