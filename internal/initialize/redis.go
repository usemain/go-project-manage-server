package initialize

import (
	"fmt"
	"gin-choes-server/internal/global"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

func InitRedis() {
	DSN := fmt.Sprintf("127.0.0.1:6379")
	global.GVA_REDIS = redis.NewClient(&redis.Options{
		Addr: DSN,
	})
}
