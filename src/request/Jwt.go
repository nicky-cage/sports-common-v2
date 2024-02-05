package request

import (
	"fmt"
	"sports-common/consts"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetJwtToken 反解析token
func GetJwtToken(c *gin.Context) jwt.MapClaims {
	authString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(consts.JwtSecret), nil
	})
	if err != nil {
		return nil
	}

	return token.Claims.(jwt.MapClaims)
}
