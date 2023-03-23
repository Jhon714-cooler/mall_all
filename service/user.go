package service

import (
	"context"
	"log"
	"mall/dao"
	"mall/middleware"
	"mall/models"
	"mall/respond"
	"mall/serialize"
)

type UserService struct {
	Username string `form:"user_name" json:"user_name"`
	Nickname string `form:"nick_name" json:"nick_name"`
	Passwd   string `form:"passwd" json:"passwd"`
}
type SendEmail struct{
	Email string `json:"email" form:"email"`
	Passwd string `json:"passwd" form:"passwd"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}
//注册
func (Service UserService)Regist(c context.Context)(respond.Response)  {
	
	var user *models.User
	userDao := dao.NewUserDao(c)
	_, exist, err := userDao.ExistOrNotByUserName(Service.Username)
	if err != nil {
		code := respond.ErrorDatabase
		log.Fatalf("slect UserName err",err)
		return  respond.Response{
			Code: 500,
			Message: respond.GetMsg(code),
			Data: nil,
		}
	}
	if exist {
		code := respond.ErrorExistUser
		return respond.Response{
			Code:500,
			Message: respond.GetMsg(code),
			Data: nil,
		}
	}
	user = &models.User{
		NickName: Service.Nickname,
		UserName: Service.Username,
		PasswordDigest: models.SetPassWd(Service.Passwd),
	}
	user.Avatar = "avatar.JPG"
	//创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		return respond.Response{
			Code:500,
			Message: respond.GetMsg(40001),
			Data: nil,
		}
	}
	return respond.Response{
		Code: 200,
		Message: respond.GetMsg(200),
		Data: "create succeed！",
	}
}
func (Service UserService)Login(c context.Context)(respond.Response)  {
	var user *models.User
	userDao := dao.NewUserDao(c)
	user, exist, err := userDao.ExistOrNotByUserName(Service.Username)
	if err != nil {
		code := respond.ErrorDatabase
		return respond.Response{
			Code: 202,
			Message: respond.GetMsg(code),
			Data: nil,
		}
	}
	if !exist{
		code := respond.ErrorUserNotFound
		return respond.Response{
			Code: 202,
			Message: respond.GetMsg(code),
		}
	}
	if (user.CheckPasswd(Service.Passwd)) == false{
		code := respond.ErrorNotComparePassword
		return respond.Response{
			Code: 202,
			Message: respond.GetMsg(code),
		}
	}
	token ,err := middleware.GenerateToken(user.ID,user.UserName,0)
	if err != nil {
		code := respond.ErrorAuthToken
		return respond.Response{
			Code: 500,
			Message: respond.GetMsg(code),
		}
	}
	return respond.Response{
		Code: 200,
		Message: respond.GetMsg(200),
		Data: serialize.TokenData{User: serialize.BuildUser(user), Token: token},
	}
}