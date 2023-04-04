package router

import (
	"mall/api"
	"mall/middleware"

	"github.com/gin-gonic/gin"
)

func Router()  {
	r := gin.Default()
	r.Use(middleware.Cors())
	{
		r.POST("user/register", api.Regist)
		r.POST("user/login", api.Login)
		r.Use(middleware.JWT())


		authed := r.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)
			authed.POST("avatar", api.UploadAvatar)
			
		}
	}
	r.Run(":8081")
}