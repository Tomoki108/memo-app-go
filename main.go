package main

import (
	"fmt"
	"memo-app-go/database"
	"time"

	"github.com/gin-gonic/gin"
)

type Memo struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	r := gin.Default()

	r.GET("/memos", func(c *gin.Context) {
		var memos []Memo
		database.Db.Find(&memos)
		c.JSON(200, memos)
	})

	r.POST("/memos", func(c *gin.Context) {
		var memo Memo
		if err := c.ShouldBindJSON(&memo); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(400)
		} else if memo.Title == "" {
			fmt.Print("title is empty yo")
			c.AbortWithStatus(400)
		} else {
			database.Db.Create(&memo)
			c.JSON(200, memo)
		}
	})

	r.GET("/memos/:id", func(c *gin.Context) {
		var memo Memo
		if err := database.Db.Where("id = ?", c.Param("id")).First(&memo).Error; err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
		} else {
			c.JSON(200, memo)
		}
	})

	r.PUT("/memos/:id", func(c *gin.Context) {
		var memo Memo
		if err := database.Db.Where("id = ?", c.Param("id")).First(&memo).Error; err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
		} else {
			if err := c.ShouldBindJSON(&memo); err != nil {
				c.AbortWithStatus(400)
			} else {
				database.Db.Save(&memo)
				c.JSON(200, memo)
			}
		}
	})

	r.DELETE("/memos/:id", func(c *gin.Context) {
		var memo Memo
		if err := database.Db.Where("id = ?", c.Param("id")).First(&memo).Error; err != nil {
			fmt.Println(err)
			c.AbortWithStatus(404)
		} else {
			database.Db.Delete(&memo)
			c.JSON(200, gin.H{"id #" + c.Param("id"): "deleted"})
		}
	})

	r.Run() // defaultの8080ポートでAPIを公開
}
