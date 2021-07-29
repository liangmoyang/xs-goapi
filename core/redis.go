package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
)

func NewRedis(ctx context.Context) *redis.Client {

	host := AppConfig.GetString("redis.host")
	port := AppConfig.GetString("redis.port")
	password := AppConfig.GetString("redis.password")
	db := AppConfig.GetInt("redis.db")

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	//defer client.Close()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(errors.New("Redis 启动失败:" + err.Error()))
	}

	log.Println("Redis连接成功")

	return client
}
