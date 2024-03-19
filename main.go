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
	cmd.Run(":8888")
}
