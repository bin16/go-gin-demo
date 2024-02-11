package models

import "time"

type Model struct {
	ID        int64     `json:"id" grom:"primaryKey,unique,autoIncrement" uri:"id"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"-" gorm:"autoDeleteTime"`
}

// page query
type Query struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}
