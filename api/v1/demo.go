package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func F1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "请求成功"})
}
