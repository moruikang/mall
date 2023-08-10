package initialize

import (
	"github.com/go-redis/redis/v8"
	"mall.com/config/global"
)

func Redis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Auth, // no password set
		DB:       r.Db,   // use default DB
	})
	global.Rdb = rdb
}
