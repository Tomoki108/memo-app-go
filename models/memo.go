package models

import "time"

type Memo struct {
	Id        int       `json:"id" gorm:"primary_key" extensions:"x-order=0"`
	Title     string    `json:"title" gorm:"not null" extensions:"x-order=1"`
	Body      string    `json:"body" extensions:"x-order=2"`
	CreatedAt time.Time `json:"created_at" extensions:"x-order=3"`
	UpdatedAt time.Time `json:"updated_at" extensions:"x-order=4"`
}
