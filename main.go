package main

import (

	"mall/initialize"
)
func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	//initialize.Elastic()
	initialize.Redis()
	initialize.Init_router()
}
