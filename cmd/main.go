package main

import (
	"go-project-manage-server/internal/global"
	"go-project-manage-server/internal/initialize"
	_ "go-project-manage-server/internal/logic"
	"go-project-manage-server/internal/model"
	"log"
)

func init() {
	initialize.InitRedis()
	initialize.InitMySQL()
}

func main() {
	models := []interface{}{
		&model.TASK{},
		&model.USER{},
		&model.GROUP{},
	}
	for _, item := range models {
		if err := global.GVA_DB.AutoMigrate(item); err != nil {
			log.Panicf("auto migrate failed: %v\n", err)
		}
	}
	Run(":8888")
}
