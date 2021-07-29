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
			}
		}()

		if core.Config.App.Debounce == 1 && core.Global.Redis != nil {
			// 后端防抖
			ip := c.Request.Header.Get("X-real-ip")
			if ip == "" {
				ip = c.ClientIP()
			}

			key := ip + c.Request.URL.Path
			_, err := core.Global.Redis.Get(c, key).Result()
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"code": http.StatusForbidden, "msg": "重复请求"})
				c.Abort()
			} else {
				if core.Global.Redis.Set(c, key, time.Now(), 200*time.Millisecond).Err() != nil {
					fmt.Println("MiddleWare - app.go - Redis设置值失败，防抖失效")
				}
			}
		}

		c.Next()
	}
}
