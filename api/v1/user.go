package v1

import (
	"github.com/gin-gonic/gin"
	"goWebDemo/api/v1/validate"
	"goWebDemo/model"
	"goWebDemo/service"
	"goWebDemo/utils/errmsg"
	"goWebDemo/utils/response"
	"goWebDemo/utils/validator"
	"strconv"
)

// AddUser 新增用戶
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.Success {
		response.Fail(c, validCode, msg, "")
		return
	}
	code := service.CheckUser(data.Username)
	if code == errmsg.Success {
		service.CreateUser(&data)
		response.Success(c, "创建用户成功", "")
	} else {
		response.Fail(c, code, errmsg.GetErrMsg(code), "")
		return
	}
}

// GetUser 获取用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.GetUser(id)
	if code != errmsg.Success {
		response.Fail(c, code, errmsg.GetErrMsg(code), "")
		return
	}
	response.Success(c, "获取用户成功", gin.H{"data": data})
}

// GetUserList 获取用户列表
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
	response.Success(c, "获取用户列表成功", gin.H{"data": data, "total": total})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.Fail(c, errmsg.Error, err.Error(), "")
	}

	code := service.CheckUpUser(id, data.Username)
	if code != errmsg.Success {
		response.Fail(c, code, errmsg.GetErrMsg(code), "")
		return
	}
	service.EditUser(id, &data)
	response.Success(c, "更新用户成功", "")

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := service.DeleteUser(id)
	if code != errmsg.Success {
		response.Fail(c, code, errmsg.GetErrMsg(code), "")
		return
	}
	response.Success(c, "删除用户成功", "")

}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var params validate.Password
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&params)
	msg, validCode := validator.Validate(&params)
	if validCode != errmsg.Success {
		response.ErrorParam(c, validCode, msg, "")
		return
	}

	code := service.ChangePassword(id, params.Password)
	if code != errmsg.Success {
		response.Fail(c, code, "修改密码失败", "")
		return
	}
	response.Success(c, "修改密码成功", "")
}
