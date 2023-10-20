package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"memo-app-go/router"
)

type Memo struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	engine := gin.Default()
	router.SetRoutes(engine)
	engine.Run() // defaultの8080ポートでAPIを公開
}
