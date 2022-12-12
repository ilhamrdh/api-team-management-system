package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EntityException(c *gin.Context, message string) {
	res := Error{
		Success: false,
		Message: message,
		Data:    nil,
	}

	c.JSON(http.StatusUnprocessableEntity, res)
}
