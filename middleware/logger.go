package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Log() gin.HandlerFunc {
	filePath := "log/log"
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("error: ", err)
	}
	logger := logrus.New()
	logger.Out = f
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 每一周清除日志
		rotatelogs.WithRotationTime(24*time.Hour), // 按照每天分割日志
	)

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.TraceLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Since(startTime).Milliseconds()
		spendTime := fmt.Sprintf("%d ms", endTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		RequestURI := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":   hostName,
			"RequestUrl": RequestURI,
			"Method":     method,
			"StatusCode": statusCode,
			"SpendTime":  spendTime,
			"ClientIp":   clientIp,
			"User-Agent": userAgent,
			"DataSize":   dataSize,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode == 500 {
			entry.Error("Unknown Error")
		} else if statusCode == 404 {
			entry.Warn("Page Not Found")
		} else {
			entry.Info()
		}
	}
}
