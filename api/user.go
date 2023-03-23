package api

import (
	"fmt"
	"mall/respond"
	"mall/service"

	"github.com/gin-gonic/gin"
)

func Regist(c *gin.Context)  {
	var userRegisterService service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Regist(c.Request.Context())
		c.JSON(200, res)
	}// else {
	// 	c.JSON(400, ErrorResponse(err))
	// 	util.LogrusObj.Infoln(err)
	// }
}
func Login(c *gin.Context)  {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, respond.Response{
			Code: respond.ERROR,
			Error: fmt.Sprint(err),
		})
	}
}