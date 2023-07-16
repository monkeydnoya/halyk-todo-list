package database

import (
	"context"
	"errors"
	"halyk-todo-list-api/internal/domain/models"
)

type Database interface {
	Start() error

	GetTask(ctx context.Context, id string) (models.Task, error)
	GetTasks(ctx context.Context, filter models.Filter) ([]models.Task, error)

	CreateTask(ctx context.Context, task models.CreateTask, tags []models.Tag) (models.Task, error)
	UpdateTask(ctx context.Context, id string, task models.UpdateTask) error

	AddCheckListItemToTask(ctx context.Context, task models.Task, newCheckList models.CreateCheckList) error
	AddTagToTask(ctx context.Context, task models.Task, tag models.Tag) error

	CompleteTask(ctx context.Context, id string) error

	DeleteCheckListItemFromTask(ctx context.Context, task models.Task, checkList models.CheckList) error
	DeleteTagFromTask(ctx context.Context, task models.Task, tag models.Tag) error
	DeleteTask(ctx context.Context, task models.Task) error

	GetCheckListItem(ctx context.Context, id string) (models.CheckList, error)
	CompleteCheckListItem(ctx context.Context, id string) error

	GetTag(ctx context.Context, id string) (models.Tag, error)
	GetTags(ctx context.Context) ([]models.Tag, error)
	CreateTag(ctx context.Context, tag models.CreateTag) (models.Tag, error)
	UpdateTag(ctx context.Context, id string, tag models.UpdateTag) error
	DeleteTag(ctx context.Context, tag models.Tag) error
}

var (
	ErrorUDTUnavailable  = errors.New("UDT unavailable")
	ErrUnknownRetryType  = errors.New("unknown retry type")
	ErrUnsupported       = errors.New("unsupported")
	ErrUnknownError      = errors.New("unknown error")
	ErrNotFound          = errors.New("not found")
	ErrAlreadyExists     = errors.New("already exist")
	ErrUnsupportedStatus = errors.New("unsupported status type")
	ErrDuplicatedKey     = errors.New("duplication key error")
	ErrPrimaryKey        = errors.New("primary key is not set")
	ErrInvalidField      = errors.New("object has invalid field")
	ErrInvalidData       = errors.New("invalid data object")
	ErrInvalidTx         = errors.New("invalid transaction")
)
