package main

import (
	"context"
	"goapi/core"
	"goapi/model"
	"goapi/router"
	"log"
)

func main() {

	// 启动mongo
	mgCtx, cancelMg := context.WithCancel(context.Background())
	defer cancelMg()
	core.Global.Mongo = core.NewMongoLog(mgCtx)

	// 启动Redis
	rdsCtx, cancelRds := context.WithCancel(context.Background())
	defer cancelRds()
	core.Global.Redis = core.NewRedis(rdsCtx)

	// 程序结束前关闭数据库连接
	db, _ := model.Db.DB()
	defer db.Close()

	r := router.Router()

	log.Fatalln(r.Run(":9595"))

}
