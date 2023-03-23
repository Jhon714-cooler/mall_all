package dao

import (
	"context"
	"mall/global"
	"mall/models"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}
func NewUserDao(ctx context.Context) (*UserDao) {
	return &UserDao{NewDBClient(ctx)}
}
// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao)ExistOrNotByUserName(username string) (user *models.User, exist bool, err error) {
	var count int64
	err = global.Db.Model(&models.User{}).Where("user_name=?", username).Count((&count)).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&models.User{}).Where("user_name=?", username).First((&user)).Error

	return user, true, err
}
// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *models.User) error {
	return dao.DB.Model(&models.User{}).Create(&user).Error
}

