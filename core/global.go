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
	App      app
	DataBase database
	Mongo    mg
	Redis    rds
}

// App的应用配置
type app struct {
	Debounce int
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
	Host  string
	Port  string
	LogDb string
}

// Redis
type rds struct {
	Host     string
	Port     string
	Password string
	DB       int
}
