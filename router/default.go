package router

import (
	v1 "easy-admin/api/v1"
	"easy-admin/middleware"
	"easy-admin/utils"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	gin.SetMode(utils.AppMode)
	r = gin.New()
	r.Use(middleware.Log)
	r.Use(middleware.Recovery)
	registry()
}

func registry() error {
	v1R := r.Group("/api/v1")
	{
		v1R.GET("ping", v1.Ping)
		v1R.POST("user", v1.UserAdd)
		v1R.GET("user/:id", v1.UserInfo)
		v1R.GET("user", v1.UserList)
		v1R.DELETE("user/:id", v1.UserDel)
	}
	return nil
}

func Start() {
	r.Run(utils.AppAddress)
}
