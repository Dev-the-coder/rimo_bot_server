package router

import (
	"rimo_bot_server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes and map them to handlers
	// r.POST("/bot", handlers.HomeHandler)
	r.GET("/hello", handlers.HelloHandler)
	r.GET("/json", handlers.JSONGetHandler)
	r.POST("/json", handlers.JSONPostHandler)

	r.POST("/bot", handlers.BotHandler)

	return r
}
