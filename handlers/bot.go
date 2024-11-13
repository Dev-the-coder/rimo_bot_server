// handlers/bot.go
package handlers

import (
	"net/http"
	"rimo_bot_server/services"
	"rimo_bot_server/utils"

	"github.com/gin-gonic/gin"
)

func BotHandler(c *gin.Context) {
	// Ensure the request method is POST
	if c.Request.Method != http.MethodPost {
		c.Header("Allow", http.MethodPost)
		utils.JSONErrorResponse(c, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Generate UUID
	uuidStr, err := services.GenerateUUID()
	if err != nil {
		utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to generate UUID")
		return
	}

	// Prepare and send response
	response := gin.H{
		"uuid": uuidStr,
	}
	utils.JSONResponse(c, http.StatusOK, response)
}
