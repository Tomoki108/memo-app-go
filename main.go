package main

import (
	"time"
	"you-tree-backend/database"

	"github.com/gin-gonic/gin"
)

type Tree struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	var tree Tree
	database.Db.First(&tree, 1)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, tree)
	})
	r.Run()
}
