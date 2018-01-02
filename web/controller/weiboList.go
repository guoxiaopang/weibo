package controller

import (
	"github.com/gin-gonic/gin"
)

// WeiboListController 微博列表控制器
type WeiboListController struct{}

// RegisterRoute 注册路由
func (c WeiboListController) RegisterRoute(e *gin.Engine) {
	e.GET("/list", test)
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "测试",
		"nick":    "nickName",
	})
}
