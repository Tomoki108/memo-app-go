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

// @Summary Get a memo
// @Description Get a memo
// @ID get-memo
// @Accept  json
// @Produce  json
// @Param id path int true "Memo ID"
// @Success 200 {object} models.Memo
// @Failure 400
// @Failure 404
// @Router /memos/{id} [get]
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

// @Summary Create a memo
// @Description Create a memo
// @ID post-memo
// @Accept  json
// @Produce  json
// @Param memo body models.Memo true "Memo"
// @Success 201 {object} models.Memo
// @Failure 400
// @Failure 500
// @Router /memos [post]
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

// @Summary Update a memo
// @Description Update a memo
// @ID put-memo
// @Accept  json
// @Produce  json
// @Param id path int true "Memo ID"
// @Param memo body models.Memo true "Memo"
// @Success 200 {object} models.Memo
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /memos/{id} [put]
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

// @Summary Delete a memo
// @Description Delete a memo
// @ID delete-memo
// @Accept  json
// @Produce  json
// @Param id path int true "Memo ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /memos/{id} [delete]
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

// @Summary Get memos
// @Description Get memos
// @ID get-memos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Memo
// @Failure 500
// @Router /memos [get]
func GetMemos(c *gin.Context) {
	var memos []models.Memo
	if err := database.Db.Find(&memos).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, memos)
}
