package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestException(c *gin.Context, message string) {
	res := Error{
		Success: false,
		Message: message,
		Data:    nil,
	}

	c.JSON(http.StatusBadRequest, res)
}
