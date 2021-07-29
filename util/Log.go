package util

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/core"
	"os"
	"runtime"
	"strings"
	"time"
)

func ErrorFormatString(err interface{}) string {
	var text string

	//fmt.Printf("%T", err)

	switch e := err.(type) {
	case *runtime.TypeAssertionError:
		text = "参数类型不正确，导致发生错误。" + e.Error()
	case runtime.Error:
		text = e.Error()
	case error:
		text = e.Error()
	default:
		text = e.(string)
	}

	return text
}

// AppLog 写入App日志到Mongo
func AppLog(logText ...string) {

	collection := core.Global.Mongo.Collection("app_log")

	appLog := new(appLog)
	appLog.Time = time.Now().Format("2006-01-02 15:04:05")

	for _, v := range logText {
		appLog.Content = appLog.Content + v + ";"
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, &appLog)
	if err != nil {
		fmt.Println("Mongo App Log Save Failed")
	}

	return
}

type appLog struct {
	Time    string
	Content string
}
type errLog struct {
	Time     string
	Host     string
	FileInfo string
	Line     int
	FileName string
	ErrMsg   string
}

// ErrLog 记录错误日志到Mongo
func ErrLog(c *gin.Context, logText string) (err error) {
	collection := core.Global.Mongo.Collection("err_log")

	pc, fileInfo, line, _ := runtime.Caller(3)
	f := runtime.FuncForPC(pc)

	errLog := new(errLog)
	errLog.Time = time.Now().Format("2006-01-02 15:04:05")
	errLog.Host = c.Request.Host
	errLog.FileInfo = fileInfo
	errLog.Line = line
	errLog.FileName = f.Name()
	errLog.ErrMsg = logText

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, &errLog)
	if err != nil {
		fmt.Println("Mongo Err Log Save Failed")
		return err
	}

	return
}

// Log 写日志到app_log目录
// 提供给Primary类型日志使用
func Log(logText ...string) (err error) {

	path := "log/app_log/"

	if !Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	fileName := path + time.Now().Format("2006-01-02") + ".txt"

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	var builder strings.Builder
	builder.WriteString("[" + time.Now().Format("2006-01-02 15:04:05") + "]\t" + "\t")

	for _, v := range logText {
		builder.WriteString(v + "\t")
	}

	_, err = file.WriteString(builder.String())
	if err != nil {
		return err
	}
	return
}
