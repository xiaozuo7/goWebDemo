package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	v1 "goWebDemo/api/v1"
	"goWebDemo/middleware"
)

func InitRouter() {
	gin.SetMode(viper.GetString("Server.AppMode"))
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)
		auth.PUT("/user/changepswd/:id", v1.ChangePassword)

	}
	router := r.Group("api/v1")
	{
		router.POST("/login", v1.Login)
		router.POST("/upload", v1.Upload)
		router.GET("/user/:id", v1.GetUser)
		router.GET("/users", v1.GetUserList)
		router.POST("/user/add", v1.AddUser)
		router.POST("/front/login", v1.LoginFront)

	}

	_ = r.Run(viper.GetString("Server.HttpPort"))

}
