package initialize

import (
	"fmt"
	"gin_skeleton/g"
	"github.com/go-redis/redis"
	"time"
)

// 初始化redis
func InitRedis() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:        g.Config.GetString("redis.addr"),
		Password:    g.Config.GetString("redis.password"),
		DB:          0,
		PoolSize:    50,
		PoolTimeout: time.Second * time.Duration(600),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	g.Redis = redisClient
}
