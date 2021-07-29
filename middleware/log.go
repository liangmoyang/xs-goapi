package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/core"
	"io/ioutil"
	"time"
)

type log struct {
	Time      string
	Ip        string
	Status    int
	Method    string
	Host      string
	Path      string
	PostParam string
	GetParam  string
}

func Log() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取post的body
		buf := make([]byte, 1024)
		n, _ := ctx.Request.Body.Read(buf)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))

		// 获取get的param
		var getJson []byte
		query := ctx.Request.URL.Query()
		if len(query) != 0 {
			getJson, _ = json.Marshal(ctx.Request.URL.Query())
		}

		ctx.Next()

		collection := core.Global.Mongo.Collection("api_log")

		_, err := collection.InsertOne(ctx, log{
			time.Now().Format("2006-01-02 15:04:05"),
			ctx.Request.Header.Get("X-real-ip"),
			ctx.Writer.Status(),
			ctx.Request.Method,
			ctx.Request.Host,
			ctx.Request.URL.Path,
			string(buf[0:n]),
			string(getJson),
		})
		if err != nil {
			fmt.Println("Mongo日志保存失败")
		}
	}
}
