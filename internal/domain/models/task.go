package models

type UpdateTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline,omitempty"`
}

type CreateTask struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Deadline    int64             `json:"deadline,omitempty"`
	CheckList   []CreateCheckList `json:"checklist,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
}

type Task struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Deadline    int64       `json:"deadline,omitempty"`
	CreatedAt   int64       `json:"created_at"`
	FinishedAt  int64       `json:"finished_at"`
	Done        bool        `json:"done"`
	CheckList   []CheckList `json:"checklist"`
	Tags        []Tag       `json:"tags"`
}

type CreateCheckList struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CheckList struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FinishedAt  int64  `json:"finished_at"`
	Done        bool   `json:"done"`
}
