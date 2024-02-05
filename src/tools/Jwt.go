package tools

import (
	"time"

	"sports-common/consts"

	"github.com/dgrijalva/jwt-go"
)

// GetJwtToken 生成token
func GetJwtToken(id uint64, username string, roleId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id                                    //用户编号
	claims["name"] = username                            //用户名称
	claims["role_id"] = roleId                           //角色编号
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //过期时间
	claims["iat"] = time.Now().Unix()
	claims["sub"] = id
	token.Claims = claims

	return token.SignedString([]byte(consts.JwtSecret))
}
