package models

type NoteVisible string

const (
	PUBLIC_NOTE  NoteVisible = "public"
	PRIVATE_NOTE NoteVisible = "private"
)

type Note struct {
	Model
	Content   string      `json:"content"`
	Visible   NoteVisible `json:"visible" gorm:"default:private"`
	ProfileID int64       `json:"profileId"`
	Profile   Profile     `json:"profile"`
}
