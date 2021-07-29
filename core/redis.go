package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
)

func NewRedis(ctx context.Context) *redis.Client {

	host := Config.Redis.Host
	port := Config.Redis.Port
	password := Config.Redis.Password
	db := Config.Redis.DB

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
