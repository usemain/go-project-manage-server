package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"go-project-manage-server/internal/global"
	"go-project-manage-server/internal/model"
	"gopkg.in/yaml.v3"
	"os"
)

func InitRedis() {
	var ds = &model.DATABASES{}
	filePath := "config/databases.yaml"
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read file -> " + err.Error())
	}
	if err = yaml.Unmarshal(file, ds); err != nil {
		panic("Yaml unmarshal error -> " + err.Error())
	}

	DSN := fmt.Sprintf("%s:%v", ds.REDIS.Host, ds.REDIS.Port)
	global.GVA_REDIS = redis.NewClient(&redis.Options{
		Addr: DSN,
	})
}
