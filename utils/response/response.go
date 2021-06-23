package response

import (
	"github.com/gin-gonic/gin"
	"goWebDemo/utils/errmsg"
	"net/http"
)

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	c.JSON(httpCode, gin.H{
		"status": dataCode,
		"msg":    msg,
		"data":   data,
	})
}

// Success 返回成功 200
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, errmsg.Success, msg, data)
}

// Fail 返回失败 400
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// ErrorSystem 服务器错误 500
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, errmsg.Error, msg, data)
	c.Abort()
}

// ErrorParam 参数校验失败 400
func ErrorParam(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}
