package models

type Filter struct {
	Done        string `url:"done,omitempty"`
	CreatedFrom int64  `url:"created_from,omitempty"`
	CreatedTo   int64  `url:"created_to,omitempty"`
}
