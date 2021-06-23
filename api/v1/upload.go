package v1

import (
	"github.com/gin-gonic/gin"
	"goWebDemo/service"
	"goWebDemo/utils/errmsg"
	"goWebDemo/utils/response"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil || fileHeader == nil {
		response.ErrorParam(c, errmsg.Error, "参数校验失败", "")
		return
	}
	fileSize := fileHeader.Size
	url, code := service.UpLoadFile(file, fileSize)
	if code != errmsg.Success {
		response.Fail(c, code, errmsg.GetErrMsg(code), "")
		return
	}
	response.Success(c, "上传文件成功", gin.H{"url": url})
}
