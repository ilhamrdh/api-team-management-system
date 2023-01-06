package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundException(c *gin.Context, message string) {
	res := Error{
		Success: false,
		Message: message,
		Data:    nil,
	}

	c.JSON(http.StatusNotFound, res)
}
