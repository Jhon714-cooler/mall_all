package initialize

import (
	"mall/dao"
	"mall/global"
	"strings"
)

func mysql_init()  {
		//MySQL
		pathRead := strings.Join([]string{global.Config.Mysql.Username, ":", global.Config.Mysql.Password,"@",global.Config.Mysql.Url},"")
		pathWrite := strings.Join([]string{global.Config.Mysql.Username, ":", global.Config.Mysql.Password,"@",global.Config.Mysql.Url},"")
		dao.Database(pathRead, pathWrite)
}