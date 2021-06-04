package v1

import (
	"github.com/gin-gonic/gin"
	"goWebDemo/api/v1/validate"
	"goWebDemo/model"
	"goWebDemo/service"
	"goWebDemo/utils/errmsg"
	"goWebDemo/utils/validator"
	"net/http"
	"strconv"
)

var code int

// AddUser 新增用戶
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.Success {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code = service.CheckUser(data.Username)
	if code == errmsg.Success {
		service.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUser 获取用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.GetUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10

	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := service.GetUserList(username, pageSize, pageNum)

	code = errmsg.Success
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
}

func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = service.CheckUpUser(id, data.Username)
	if code == errmsg.Success {
		service.EditUser(id, &data)
	}
	if code == errmsg.ErrorUserNameExists {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = service.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var params validate.Password
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&params)
	msg, validCode := validator.Validate(&params)
	if validCode != errmsg.Success {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code = service.ChangePassword(id, params.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
