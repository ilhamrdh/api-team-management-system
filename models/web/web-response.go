package web

import (
	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

func ResponseSuccess(c *gin.Context, data interface{}, message string, status int) {
	res := WebResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	c.JSON(status, res)
}
