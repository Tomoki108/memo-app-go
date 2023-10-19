package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"memo-app-go/database"
	"memo-app-go/models"

	"github.com/gin-gonic/gin"
)

func GetMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	var memo models.Memo
	if err := database.Db.First(&memo, id).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
		return
	}

	c.JSON(http.StatusOK, memo)
}

func PostMemo(c *gin.Context) {
	var memo models.Memo
	if err := c.ShouldBindJSON(&memo); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	if memo.Title == "" {
		fmt.Println("title is empty yo")
		c.AbortWithStatus(400)
		return
	}

	memo.CreatedAt = time.Now()
	memo.UpdatedAt = time.Now()

	if err := database.Db.Create(&memo).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, memo)
}

func PutMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	var memo models.Memo
	if err := database.Db.First(&memo, id).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
		return
	}

	if err := c.ShouldBindJSON(&memo); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	if memo.Title == "" {
		fmt.Println("title is empty yo")
		c.AbortWithStatus(400)
		return
	}

	memo.UpdatedAt = time.Now()

	if err := database.Db.Save(&memo).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, memo)
}

func DeleteMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	var memo models.Memo
	if err := database.Db.First(&memo, id).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
		return
	}

	if err := database.Db.Delete(&memo).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.Status(http.StatusNoContent)
}

func GetMemos(c *gin.Context) {
	var memos []models.Memo
	if err := database.Db.Find(&memos).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, memos)
}
