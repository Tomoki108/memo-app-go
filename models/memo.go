package models

import "time"

type Memo struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
