package entities

import (
	uuid "github.com/satori/go.uuid"
)

type Task struct {
	Id          uuid.UUID   `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Deadline    int64       `json:"deadline,omitempty"`
	CreatedAt   int64       `json:"created_at"`
	FinishedAt  int64       `json:"finished_at"`
	Done        bool        `json:"done"`
	CheckList   []CheckList `json:"checklists" gorm:"foreignKey:TaskId"`
	Tags        []Tag       `json:"tags" gorm:"many2many:task_tags;"`
}

func (t *Task) TableName() string {
	return "tasks"
}

type CheckList struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FinishedAt  int64     `json:"finished_at"`
	Done        bool      `json:"done"`
	TaskId      uuid.UUID `json:"task_id"`
}

func (cl *CheckList) TableName() string {
	return "check_lists"
}
