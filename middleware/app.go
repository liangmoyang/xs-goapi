package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/core"
	"goapi/util"
	"net/http"
	"time"
)

func App() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {

				fmt.Println(err) // 打印出来方便查看

				errString := util.ErrorFormatString(err)

				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "请求失败：" + errString,
				})

				_ = util.ErrLog(c, errString)

				// 发送给群机器人
				go wxRobot(errString)
			}
		}()

		// 后端防抖
		ip := c.Request.Header.Get("X-real-ip")
		if ip == "" {
			ip = c.ClientIP()
		}

		key := ip + c.Request.URL.Path
		_, err := core.Global.Redis.Get(c, key).Result()
		if err == nil {
			c.JSON(200, gin.H{"code": http.StatusForbidden, "msg": "重复请求"})
			c.Abort()
		} else {
			err := core.Global.Redis.Set(c, key, time.Now(), 200*time.Millisecond).Err()
			if err != nil {
				fmt.Println("MiddleWare - app.go - Redis设置值失败，防抖失效")
			}
		}

		c.Next()
	}
}

func wxRobot(text string) {

	if err := recover(); err != nil {
		fmt.Println(err)
	}

	webHook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=300bdddb-1333-4ce2-8f89-e93599d1ab5c"

	content := make(map[string]interface{})
	content["content"] = text

	param := struct {
		MsgType string                 `json:"msgtype,omitempty"`
		Text    map[string]interface{} `json:"text,omitempty"`
	}{
		"text", content,
	}

	_, _ = util.PostStruct(webHook, param)
}
