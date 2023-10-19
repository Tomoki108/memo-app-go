package router

import (
	"memo-app-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	r.GET("/memos", handlers.GetMemos)
	r.POST("/memos", handlers.PostMemo)
	r.GET("/memos/:id", handlers.GetMemo)
	r.PUT("/memos/:id", handlers.PutMemo)
	r.DELETE("/memos/:id", handlers.DeleteMemo)
}
