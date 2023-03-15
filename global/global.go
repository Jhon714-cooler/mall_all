package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"mall/conf"
)

var (
	Config conf.Config
	Db     *gorm.DB
	Rdb    *redis.Client
	Es     *elastic.Client
)
