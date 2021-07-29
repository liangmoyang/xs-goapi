package core

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	mongoDB "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func NewMongoLog(ctx context.Context) *mongoDB.Database {

	host := Config.Mongo.Host
	port := Config.Mongo.Port
	dbname := Config.Mongo.LogDb

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	client, err := mongoDB.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(errors.New("MongoDB连接失败：" + err.Error()))
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(errors.New("MongoDB连接接失败：" + err.Error()))
	}

	db := client.Database(dbname)
	log.Println("MongoDB连接成功")

	return db
}
