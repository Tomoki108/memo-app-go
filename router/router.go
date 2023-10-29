package router

import (
	"memo-app-go/docs"
	"memo-app-go/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRoutes(engine *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "memo app"

	{
		rg := engine.Group("/api/v1")
		rg.GET("/memos", handlers.GetMemos)
		rg.POST("/memos", handlers.PostMemo)
		rg.GET("/memos/:id", handlers.GetMemo)
		rg.PUT("/memos/:id", handlers.PutMemo)
		rg.DELETE("/memos/:id", handlers.DeleteMemo)
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
