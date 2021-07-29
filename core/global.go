package core

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type global struct {
	Mongo  *mongo.Database
	Redis  *redis.Client
	Config *viper.Viper
}

var Global global
