package router

import (
	v1 "goapi/api/v1"
	"goapi/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.New()

	router.Use(middleware.Cors())
	router.Use(middleware.Log())
	router.Use(middleware.App())

	// 注册【无前缀】接口
	Group := router.Group("")
	{
		Group.GET("demo", v1.F1)
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "welcome to xs api.",
		})
	})

	// 用于健康监测
	router.GET("/h", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"code": 200})
	})

	return router

}
