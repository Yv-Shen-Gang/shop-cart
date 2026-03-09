package init

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"product-srv/config"
)

func RedisInit() {
	data := config.Config.Redis
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     data.Addr,
		Password: data.Password, // no password set
		DB:       data.Database, // use default DB
	})
	err := config.Rdb.Ping(config.Ctx).Err()
	if err != nil {
		panic("redis连接失败")
	}
	fmt.Println("redis连接成功")
}
