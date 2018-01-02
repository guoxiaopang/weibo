package main

import (
	"weibo/web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.RegisterRoutes(r)
	r.Run()
}
