package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

//func main() {
//	gin.ForceConsoleColor()		// 此方法可以在控制台显示不同颜色的log日志信息
//
//	router:=gin.Default()
//	router.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK,"Hello World")
//	})
//	router.Run(":8080")
//}

func main() {
	currenTime := time.Now()
	ct:=currenTime.Format("2006-01-02")

	_, err := os.Stat("C:\\healthcare\\"+ct)
	if err == nil {
		fmt.Println("文件夹已存在")
	}else {
		ctDir := os.Mkdir("C:\\healthcare\\"+ct, 0666)
		fmt.Println(&ctDir)
	}
}