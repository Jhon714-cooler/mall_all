package global

import (
	"github.com/go-redis/redis"
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
