package core

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func NewMongoLog(ctx context.Context) *mongo.Database {

	host := Global.Config.Get("mongo.host")
	port := Global.Config.Get("mongo.port")
	dbname := Global.Config.Get("mongo.log_dbname")

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(errors.New("MongoDB连接失败：" + err.Error()))
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(errors.New("MongoDB连接接失败：" + err.Error()))
	}

	db := client.Database(dbname.(string))
	log.Println("MongoDB连接成功")

	return db
}
