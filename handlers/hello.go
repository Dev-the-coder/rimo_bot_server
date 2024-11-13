package handlers

import (
	"net/http"
	"rimo_bot_server/services"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	message := services.GetFirstMessage()
	c.String(http.StatusOK, message)
}
