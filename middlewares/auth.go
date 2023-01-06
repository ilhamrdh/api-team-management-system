package middlewares

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/exceptions"
	"github.com/IlhamRamadhan-IR/api-team-management-system/securities"
	"github.com/gin-gonic/gin"
)

func SetupAuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := securities.GetAuthentication(ctx)
		if err != nil {
			exceptions.AuthorizeException(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
