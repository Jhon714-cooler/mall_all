package models

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
}

func SetPassWd(passwd string)(string)  {
	return fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
}
func (user *User)CheckPasswd(passwd string)bool  {
	t := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
	
	return t == user.PasswordDigest
}
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}