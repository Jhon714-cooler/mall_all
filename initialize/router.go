package initialize

import "github.com/gin-gonic/gin"

func Init_router()  {
	engine := gin.Default()

	web := engine.Group("/web")
	{
		web.POST("/login")
	}
}	