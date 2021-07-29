package core

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type global struct {
	Mongo *mongo.Database
	Redis *redis.Client
}

var Global global

func (g *global) LogDB() *mongo.Database {
	return g.Mongo
}
