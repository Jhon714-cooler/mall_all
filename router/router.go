package router

import (
	"mall/api"
	"mall/middleware"

	"github.com/gin-gonic/gin"
)

func Router()  {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.Regist)
	}
}