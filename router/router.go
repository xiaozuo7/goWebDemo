package router

import (
	"github.com/gin-gonic/gin"
	"goWebDemo/middleware"
	"goWebDemo/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
}