package initialize

import (
	"log"
	"mall/global"
	"strconv"

	"github.com/go-redis/redis"
)
func Redis() {
	
	db,err := strconv.Atoi(global.Config.Redis.RedisDbName) //strconv.ParseInt
	if err != nil {
		log.Panic("RdToerr:",err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.RedisAddr,
		Password: global.Config.Redis.RedisPw, // no password set
		DB:       db,  // use default DB
	})
	rdb.Ping().Result()
	if err != nil {
		log.Panic("can not connect to redis",err)
	}
	global.Rdb = rdb
}