package securities

import (
	"fmt"

	"github.com/IlhamRamadhan-IR/api-team-management-system/models/web/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ExtractAuthToken(c *gin.Context) (*request.UserDetail, error) {
	token, err := VerifyToken(c)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		username := fmt.Sprintf("%s", claims["username"])
		authorized := fmt.Sprintf("%s", claims["authorized"])

		if err != nil {
			return nil, err
		}

		authDetail := request.UserDetail{
			Username:  username,
			Authorize: authorized,
		}

		return &authDetail, nil
	}

	return nil, err
}
