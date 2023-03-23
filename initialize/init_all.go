package initialize

import (
	"log"
	"mall/router"
)

func Init_all()  {
	loadConfig()
	mysql_init()
	//es_init()
	redis_init()
	router.Router()
	log.Println("initialize completed")
}