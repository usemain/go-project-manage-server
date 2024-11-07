package initialize

import (
	"fmt"
	"go-project-manage-server/internal/global"
	"go-project-manage-server/internal/model"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMySQL() {
	var ds = &model.DATABASES{}
	filePath := "config/databases.yaml"
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read file -> " + err.Error())
	}
	if err = yaml.Unmarshal(file, ds); err != nil {
		panic("Yaml unmarshal error -> " + err.Error())
	}

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		ds.MySQL.User, ds.MySQL.Password, ds.MySQL.Host, ds.MySQL.Port, ds.MySQL.Database, ds.MySQL.Charset,
	)
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("MySQL connect error -> " + err.Error())
	}

	global.GVA_DB = db
}
