package v1

import (
	"github.com/gin-gonic/gin"
	"goWebDemo/model"
	"goWebDemo/utils/errmsg"
	"goWebDemo/utils/validator"
	"net/http"
)

var code int

// AddUser 新增用戶
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ =  c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.Success {
		c.JSON(
			http.StatusOK, gin.H{
			"status": validCode,
			"message": msg,
			},
		)
		c.Abort()
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.Success {
		model.CreateUser(&data)
	}


	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}