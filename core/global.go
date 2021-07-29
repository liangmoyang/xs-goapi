package core

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type global struct {
	Mongo *mongo.Database
	Redis *redis.Client
	Viper *viper.Viper
}

var (
	Global global
	Config config // 应用配置
)

type config struct {
	DataBase database
	Mongo    mg
	Redis    rds
}

// MySql
type database struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
}

// Mongo
type mg struct {
	Host      string
	Port      string
	LogDbname string
}

// Redis
type rds struct {
	Host     string
	Port     string
	Password string
	DB       int
}
