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
	r := gin.Default()

	r.GET("/trees", func(c *gin.Context) {
		var trees []Tree
		database.Db.Find(&trees)
		c.JSON(200, trees)
	})

	r.POST("/trees", func(c *gin.Context) {
		var tree Tree
		c.BindJSON(&tree)
		database.Db.Create(&tree)
		c.JSON(200, tree)
	})

	r.GET("/trees/:id", func(c *gin.Context) {
		var tree Tree
		if err := database.Db.Where("id = ?", c.Param("id")).First(&tree).Error; err != nil {
			println(err)
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, tree)
		}
	})

	r.DELETE("/trees/:id", func(c *gin.Context) {
		var tree Tree
		if err := database.Db.Where("id = ?", c.Param("id")).First(&tree).Error; err != nil {
			println(err)
			c.AbortWithStatus(404)
		} else {
			database.Db.Delete(&tree)
			c.JSON(200, gin.H{"id #" + c.Param("id"): "deleted"})
		}
	})

	r.Run() // defaultの8080ポートでAPIを公開
}
