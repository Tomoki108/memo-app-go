package main

import (
	"github.com/gin-gonic/gin"

	"memo-app-go/router"
)

func main() {
	engine := gin.Default()
	router.SetRoutes(engine)
	engine.Run() // defaultの8080ポートでAPIを公開
}
