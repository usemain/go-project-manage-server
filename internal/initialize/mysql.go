package initialize

import (
	"gin-choes-server/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMySQL() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, // 慢SQL阀值
		LogLevel:      logger.Info, // 级别
		Colorful:      true,        // 色彩
	})

	DSN := "root:a123456@tcp(127.0.0.1:3306)/gin_choes?charset=utf8mb3&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("MySQL connect error -> " + err.Error())
	}

	global.GVA_DB = db
}
