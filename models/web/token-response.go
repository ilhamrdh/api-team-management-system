package web

import "github.com/gin-gonic/gin"

type ResponseAuth struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func ResponseToken(c *gin.Context, message string, token string, status int) {
	res := ResponseAuth{
		Status:  status,
		Message: message,
		Token:   token,
	}

	c.JSON(status, res)
}
