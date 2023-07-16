package models

type UpdateTag struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Tag struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
