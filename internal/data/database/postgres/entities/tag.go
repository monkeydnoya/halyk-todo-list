package entities

import (
	uuid "github.com/satori/go.uuid"
)

type Tag struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (t *Tag) TableName() string {
	return "tags"
}
