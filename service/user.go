package service

import (
	"context"
	"log"
	"mall/dao"
	"mall/global"
	"mall/middleware"
	"mall/models"
	"mall/respond"
	"mall/serialize"
	"mall/util"
	"mime/multipart"
	"strings"

	"gopkg.in/mail.v2"
)

type UserService struct {
	Username string `form:"user_name" json:"user_name"`
	Nickname string `form:"nick_name" json:"nick_name"`
	Passwd   string `form:"passwd" json:"passwd"`
}
type EmailService struct {
	Email  string `json:"email" form:"email"`
	Passwd string `json:"passwd" form:"passwd"`
	//OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OperationType uint `form:"operation_type" json:"operation_type"`
}




// 注册
func (Service UserService) Regist(c context.Context) respond.Response {

	var user *models.User
	userDao := dao.NewUserDao(c)
	_, exist, err := userDao.ExistOrNotByUserName(Service.Username)
	if err != nil {
		code := respond.ErrorDatabase
		log.Fatalf("slect UserName err", err)
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
			Data:    nil,
		}
	}
	if exist {
		code := respond.ErrorExistUser
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
			Data:    nil,
		}
	}
	user = &models.User{
		NickName:       Service.Nickname,
		UserName:       Service.Username,
		PasswordDigest: user.SetPassWd(Service.Passwd),
	}
	user.Avatar = "avatar.JPG"
	//创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(40001),
			Data:    nil,
		}
	}
	return respond.Response{
		Code:    200,
		Message: respond.GetMsg(200),
		Data:    "create succeed！",
	}
}
func (Service UserService) Login(c context.Context) respond.Response {
	var user *models.User
	userDao := dao.NewUserDao(c)
	user, exist, err := userDao.ExistOrNotByUserName(Service.Username)
	if err != nil {
		code := respond.ErrorDatabase
		return respond.Response{
			Code:    202,
			Message: respond.GetMsg(code),
			Data:    nil,
		}
	}
	if !exist {
		code := respond.ErrorUserNotFound
		return respond.Response{
			Code:    202,
			Message: respond.GetMsg(code),
		}
	}
	if (user.CheckPasswd(Service.Passwd)) == false {
		code := respond.ErrorNotComparePassword
		return respond.Response{
			Code:    202,
			Message: respond.GetMsg(code),
		}
	}
	token, err := middleware.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		code := respond.ErrorAuthToken
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
		}
	}
	return respond.Response{
		Code:    200,
		Message: respond.GetMsg(200),
		Data:    serialize.TokenData{User: serialize.BuildUser(user), Token: token},
	}
}
func (Service UserService) UserUpdate(c context.Context, Uid uint) respond.Response {
	userdao := dao.NewUserDao(c)
	user, err := userdao.GetUserById(Uid)
	if err != nil {
		code := respond.ErrorDatabase
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
			Error:   err.Error(),
		}
	}
	if Service.Nickname != "" {
		user.NickName = Service.Nickname
	}
	if Service.Passwd != "" {
		user.PasswordDigest = user.SetPassWd(Service.Passwd)
	}
	err = userdao.UpdateByUseID(user, Uid)
	if err != nil {
		code := respond.ErrorDatabase
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return respond.Response{
		Code:    respond.SUCCESS,
		Message: "修改成功",
	}
}
func (Service UserService) UploadAvatar(c context.Context, Uid uint, file multipart.File, file_size int64) respond.Response {
	code := respond.SUCCESS
	path, err := UpLoadFile(file, file_size)
	if err != nil {
		return respond.Response{
			Code:    respond.ErrorUploadFile,
			Message: respond.GetMsg(respond.ErrorUploadFile),
			Error:   path,
		}
	}
	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserById(Uid)
	if err != nil {
		log.Println(err)
		return respond.Response{
			Code:    respond.ErrorDatabase,
			Message: respond.GetMsg(respond.ErrorDatabase),
			Error:   err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateByUseID(user, Uid)
	if err != nil {
		log.Println(err)
		return respond.Response{
			Code:    respond.ErrorDatabase,
			Message: respond.GetMsg(respond.ErrorDatabase),
			Error:   err.Error(),
		}
	}
	return respond.Response{
		Code:    code,
		Message: respond.GetMsg(code),
		Data:    serialize.BuildUser(user),
	}
}
func (Service EmailService) SendEmail(c context.Context, Uid uint) respond.Response {
	code := respond.SUCCESS
	var (
		address string
		notice  *models.Notice
	)
	token, err := util.EmailGenrateToken(Service.Email, Service.Passwd, Service.OperationType, Uid)
	if err != nil {
		code := respond.ErrorAuthToken
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
		}
	}
	noticeDao := dao.NewNoticeDao(c)
	notice, err = noticeDao.GetNoticeById(Service.OperationType)
	if err != nil {
		code := respond.ErrorDatabase
		return respond.Response{
			Code:    500,
			Message: respond.GetMsg(code),
			Error:   err.Error(),
		}
	}
	address = global.Config.Email.ValidEmail+token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", global.Config.Email.SmtpEmail)
	m.SetHeader("To", Service.Email)
	m.SetHeader("Subject", "Mall")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(global.Config.Email.SmtpHost, 465, global.Config.Email.SmtpEmail, global.Config.Email.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		code = respond.ErrorSendEmail
		return respond.Response{
			Code: code,
			Message:    respond.GetMsg(code),
		}
	}
	return respond.Response{}
}
func (Service EmailService) Valid(c context.Context, Uid uint) respond.Response {
	return respond.Response{}
}
