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



	}

	_ = r.Run(utils.HttpPort)

}
