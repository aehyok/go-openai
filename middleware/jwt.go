package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 2 // 过期时间 -2天
var Secret = []byte("i am zhaohaiyu")          // Secret(盐) 用来加密解密

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	//成员错误
	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			_, err := ParseToken(token)
			if err != nil {
				code = ErrorAuthCheckTokenFail
			}
			// todo token超期 待处理
		}

		if code != SUCCESS {
			c.JSON(400, gin.H{
				"code":    code,
				"message": "可能是身份过期了，请重新登录",
				"data":    "{}",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// 验证jwt token
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) { // 解析token
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 生成 jwt token
func GenerateToken(userID int64, username string) (string, error) {
	var claims = MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "aehyok",                                   // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", fmt.Errorf("生成token失败:%v", err)
	}
	return signedToken, nil
}
