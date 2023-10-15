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
	var trees []Tree

	database.Db.Find(&trees)

	r := gin.Default()
	r.GET("/trees", func(c *gin.Context) {
		c.JSON(200, trees)
	})

	// tree構造体のCRUD APIを実装する（GET trees以外）

	r.GET("/trees/:id", func(c *gin.Context) {
		var tree Tree
		if err := database.Db.Where("id = ?", c.Param("id")).First(&tree).Error; err != nil {
			c.AbortWithStatus(404)
			println(err)
		} else {
			c.JSON(200, tree)
		}
	})

	r.POST("/trees", func(c *gin.Context) {
		var tree Tree
		c.BindJSON(&tree)
		database.Db.Create(&tree)
		c.JSON(200, tree)
	})

	r.Run() // defaultの8080ポートでAPIを公開
}
