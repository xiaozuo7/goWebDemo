package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goWebDemo/utils/errmsg"
	"goWebDemo/utils/response"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{[]byte(viper.GetString("Server.JwtKey"))}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("token过期，请重新登录")
	TokenNotValidYet = errors.New("token无效，请重新登录")
	TokenMalformed   = errors.New("token不正确，请重新登录")
	TokenInvalid     = errors.New("请传入一个正确的token")
)

// CreateToken 生成Token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if vf, ok := err.(*jwt.ValidationError); ok {
			if vf.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if vf.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if vf.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}

// JwtToken 中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ErrorTokenExists
			response.Fail(c, code, errmsg.GetErrMsg(code), "")
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errmsg.ErrorTokenType
			response.Fail(c, code, errmsg.GetErrMsg(code), "")
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenType
			response.Fail(c, code, errmsg.GetErrMsg(code), "")
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				code = errmsg.ErrorTokenExpired
				response.Fail(c, code, errmsg.GetErrMsg(code), "")
				return
			}
			response.Fail(c, errmsg.Error, errmsg.GetErrMsg(errmsg.Error), "")
			return
		}
		c.Set("username", claims)
		c.Next()
	}

}
