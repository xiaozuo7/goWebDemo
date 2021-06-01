package router

import (
	"github.com/gin-gonic/gin"
	v1 "goWebDemo/api/v1"
	"goWebDemo/middleware"
	"goWebDemo/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())


	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())

	router := r.Group("api/v1")
	{
		router.POST("/login", v1.Login)
		router.POST("/user/add", v1.AddUser)
	}

	_ = r.Run(utils.HttpPort)

}
