package dao

import (
	"log"
	"mall/global"
	"mall/models"
	"os"
)

// Migration 执行数据迁移
func migration() {
	//自动迁移模式
	err := global.Db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&models.User{},
			&models.Product{},
			&models.Carousel{},
			&models.Category{},
			&models.Favorite{},
			&models.ProductImg{},
			&models.Order{},
			&models.Cart{},
			&models.Admin{},
			&models.Address{},
			&models.Notice{})
	if err != nil {
		log.Println("register table fail")
		os.Exit(0)
	}
	log.Println("register table success")
}
