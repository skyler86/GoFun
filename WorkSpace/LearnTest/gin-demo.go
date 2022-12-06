package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()		// 此方法可以在控制台显示不同颜色的log日志信息

	router:=gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"Hello World")
	})
	router.Run(":8080")
}

