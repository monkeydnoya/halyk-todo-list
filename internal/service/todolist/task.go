package todolist

import (
	"context"
	"halyk-todo-list-api/internal/domain/models"
)

func (s *Service) GetTask(ctx context.Context, id string) (models.Task, error) {
	task, err := s.Db.GetTask(ctx, id)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *Service) GetTasks(ctx context.Context, filter models.Filter) ([]models.Task, error) {
	tasks, err := s.Db.GetTasks(ctx, filter)
	if err != nil {
		return []models.Task{}, err
	}

	return tasks, nil
}

func (s *Service) CreateTask(ctx context.Context, task models.CreateTask) (models.Task, error) {
	tags := make([]models.Tag, 0)
	for _, taskTag := range task.Tags {
		tag, err := s.Db.GetTag(ctx, taskTag)
		if err != nil {
			return models.Task{}, err
		}
		tags = append(tags, tag)
	}
	createdTask, err := s.Db.CreateTask(ctx, task, tags)
	if err != nil {
		return models.Task{}, err
	}

	return createdTask, nil
}

func (s *Service) UpdateTask(ctx context.Context, id string, task models.UpdateTask) error {
	err := s.Db.UpdateTask(ctx, id, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddCheckListItemToTask(ctx context.Context, id string, newCheckList models.CreateCheckList) error {
	task, err := s.Db.GetTask(ctx, id)
	if err != nil {
		return err
	}

	err = s.Db.AddCheckListItemToTask(ctx, task, newCheckList)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddTagToTask(ctx context.Context, taskId string, tagId string) error {
	task, err := s.Db.GetTask(ctx, taskId)
	if err != nil {
		return err
	}
	tag, err := s.Db.GetTag(ctx, tagId)
	if err != nil {
		return err
	}

	err = s.Db.AddTagToTask(ctx, task, tag)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CompleteTask(ctx context.Context, id string) error {
	err := s.Db.CompleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteCheckListItemFromTask(ctx context.Context, taskId string, checkListId string) error {
	task, err := s.Db.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	checkList, err := s.Db.GetCheckListItem(ctx, checkListId)
	if err != nil {
		return err
	}

	err = s.Db.DeleteCheckListItemFromTask(ctx, task, checkList)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTagFromTask(ctx context.Context, taskId string, tagId string) error {
	task, err := s.Db.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	tag, err := s.Db.GetTag(ctx, tagId)
	if err != nil {
		return err
	}

	err = s.Db.DeleteTagFromTask(ctx, task, tag)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	task, err := s.Db.GetTask(ctx, id)
	if err != nil {
		return err
	}
	for _, tag := range task.Tags {
		err = s.Db.DeleteTagFromTask(ctx, task, tag)
		if err != nil {
			return err
		}
	}
	for _, checkList := range task.CheckList {
		err = s.Db.DeleteCheckListItemFromTask(ctx, task, checkList)
		if err != nil {
			return err
		}
	}
	err = s.Db.DeleteTask(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CompleteCheckListItem(ctx context.Context, id string) error {
	err := s.Db.CompleteCheckListItem(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
