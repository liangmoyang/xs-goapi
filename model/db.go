package model

import (
	"fmt"
	"goapi/core"
	"gorm.io/gorm"
	"log"
	"time"

	"gorm.io/driver/mysql"
)

var Db *gorm.DB

func InitDB() {

	username := core.Global.Config.Get("database.username")
	password := core.Global.Config.Get("database.password")
	host := core.Global.Config.Get("database.host")
	port := core.Global.Config.Get("database.port")
	dbname := core.Global.Config.Get("database.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	Db = db

	log.Println("MySql数据库初始化连接成功")
}
