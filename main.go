package main

import (
	"gin-choes-server/internal/cmd"
	"gin-choes-server/internal/initialize"
	_ "gin-choes-server/internal/logic/client"
)

func init() {
	initialize.InitRedis()
	initialize.InitMySQL()
}

func main() {
	//global.GVA_DB.AutoMigrate(model.TASK{})
	//global.GVA_DB.AutoMigrate(model.USER{})
	//global.GVA_DB.AutoMigrate(model.GROUP{})
	cmd.Run(":8888")
}
