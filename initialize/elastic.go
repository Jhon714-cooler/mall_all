package initialize

import (
	"log"
	"mall/global"

	"github.com/olivere/elastic/v7"
)

func Elastic() {
	e := global.Config.Es
	log.Println(e.EsHost+":"+e.EsPort+"/")
	client, err := elastic.NewClient(elastic.SetURL("http://"+e.EsHost+":"+e.EsPort+"/"))
	if err != nil {
		panic(err)
	}
	global.Es = client
}
