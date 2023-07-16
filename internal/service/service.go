package service

import (
	"context"
	"halyk-todo-list-api/internal/domain/models"
)

type TodoListService interface {
	GetTask(ctx context.Context, id string) (models.Task, error)
	GetTasks(ctx context.Context, filter models.Filter) ([]models.Task, error)

	CreateTask(ctx context.Context, task models.CreateTask) (models.Task, error)
	UpdateTask(ctx context.Context, id string, task models.UpdateTask) error
	CompleteTask(ctx context.Context, id string) error
	CompleteCheckListItem(ctx context.Context, id string) error

	AddCheckListItemToTask(ctx context.Context, id string, newCheckList models.CreateCheckList) error
	AddTagToTask(ctx context.Context, taskId string, tagId string) error

	DeleteCheckListItemFromTask(ctx context.Context, taskId string, checkListId string) error
	DeleteTagFromTask(ctx context.Context, taskId string, tagId string) error
	DeleteTask(ctx context.Context, id string) error

	GetTag(ctx context.Context, id string) (models.Tag, error)
	GetTags(ctx context.Context) ([]models.Tag, error)
	CreateTag(ctx context.Context, tag models.CreateTag) (models.Tag, error)
	UpdateTag(ctx context.Context, id string, tag models.UpdateTag) error
	DeleteTag(ctx context.Context, id string) error
}
