package api

import (
	"fmt"
	"mall/middleware"
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
func UserUpdate(c *gin.Context){
	var UserUpdateService service.UserService
	token := c.GetHeader("authorization")
	calims,_ := middleware.ParseToken(token)
	if err := c.ShouldBind(&UserUpdateService); err == nil{
		res := UserUpdateService.UserUpdate(c.Request.Context(),calims.ID)
		c.JSON(200, res)
	}else {
		c.JSON(400, respond.Response{
			Code: respond.ERROR,
			Error: fmt.Sprint(err),
		})
	}
}
func SendEmail(c *gin.Context){
	var EmailService service.EmailService
	token := c.GetHeader("authorization")
	calims,_ := middleware.ParseToken(token)
	if err := c.ShouldBind(&EmailService); err == nil{
		res := EmailService.SendEmail(c.Request.Context(),calims.ID)
		c.JSON(200, res)
	}else {
		c.JSON(400, respond.Response{
			Code: respond.ERROR,
			Error: fmt.Sprint(err),
		})
	}
}

func ValidEmail(c *gin.Context)  {
	var EmailService service.EmailService
	token := c.GetHeader("authorization")
	calims,_ := middleware.ParseToken(token)
	if err := c.ShouldBind(&EmailService);err == nil{
		res := EmailService.Valid(c.Request.Context(),calims.ID)
		c.JSON(200,res)
	}else{
		c.JSON(400, respond.Response{
			Code: respond.ERROR,
			Error: fmt.Sprint(err),
		})
	}
}
func UploadAvatar(c *gin.Context)  {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var UserUpdateService service.UserService
	token := c.GetHeader("authorization")
	calims,_ := middleware.ParseToken(token)
	if err := c.ShouldBind(&UserUpdateService);err == nil{
		res := UserUpdateService.UploadAvatar(c.Request.Context(),calims.ID,file,fileSize)
		c.JSON(200,res)
	}else{
		c.JSON(400, respond.Response{
			Code: respond.ERROR,
			Error: fmt.Sprint(err),
		})
	}

}