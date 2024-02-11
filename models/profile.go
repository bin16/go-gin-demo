package models

type Profile struct {
	Model
	UserID int64  `json:"-"`
	User   User   `json:"-"`
	Name   string `json:"name"`
	Notes  []Note `json:"notes,omitempty"`
}
