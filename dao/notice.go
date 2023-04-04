package dao

import (
	"context"
	"mall/models"

	"gorm.io/gorm"
)
type NoticeDao struct {
	*gorm.DB
}
func NewNoticeDao(ctx context.Context) (*NoticeDao) {
	return &NoticeDao{NewDBClient(ctx)}
}
func NewNoticedb(db gorm.DB)(*NoticeDao)  {
	return &NoticeDao{&db}
}
//修改用户信息
func (dao *NoticeDao)GetNoticeById(uId uint) (notice *models.Notice,err error) {
	err = dao.DB.Model(&models.User{}).Where("id=?", uId).First(&notice).Error
	return 
}