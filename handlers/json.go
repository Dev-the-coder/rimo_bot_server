package handlers

import (
	"net/http"
	"rimo_bot_server/services"
	"rimo_bot_server/utils"

	"github.com/gin-gonic/gin"
)

func JSONGetHandler(c *gin.Context) {
	data, err := services.GetJSONMessage()
	if err != nil {
		utils.JSONErrorResponse(c, http.StatusInternalServerError, "Failed to get JSON message")
		return
	}
	utils.JSONResponse(c, http.StatusOK, data)
}

func JSONPostHandler(c *gin.Context) {
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONErrorResponse(c, http.StatusBadRequest, "Invalid JSON")
		return
	}
	// Process input data...
	utils.JSONResponse(c, http.StatusOK, gin.H{"status": "Data received", "data": input})
}
