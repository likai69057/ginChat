package main

import (
	"ginChat/routers"
	"ginChat/utils"
)

func main() {

	utils.InitConfig()
	utils.Initmysql()

	r := routers.Router()
	r.Run(":8080")
}
