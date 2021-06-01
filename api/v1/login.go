package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goWebDemo/middleware"
	"goWebDemo/model"
	"goWebDemo/utils/errmsg"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)
	if code == errmsg.Success {
		generateToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": formData.Username,
			"id": formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token": token,
		})
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
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.Error,
			"message": errmsg.GetErrMsg(errmsg.Error),
			"token":   token,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": errmsg.Success,
		"data": user.Username,
		"id": user.ID,
		"message": errmsg.GetErrMsg(errmsg.Success),
		"token": token,
	})
	return
}
