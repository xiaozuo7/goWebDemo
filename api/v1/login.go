package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goWebDemo/middleware"
	"goWebDemo/model"
	"goWebDemo/service"
	"goWebDemo/utils/errmsg"
	"goWebDemo/utils/response"
	"time"
)

func Login(c *gin.Context) {
	var formData model.User
	var code int

	_ = c.ShouldBindJSON(&formData)
	formData, code = service.CheckLogin(formData.Username, formData.Password)
	if code == errmsg.Success {
		generateToken(c, formData)
	} else {
		response.Fail(c, code, errmsg.GetErrMsg(code), "")
		return
	}
}

func generateToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "issuer",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.ErrorSystem(c, "未知错误，请稍后再试", "")
		return
	}
	response.Success(c, "生成token成功", gin.H{"username": user.Username, "id": user.ID, "token": token})
}
