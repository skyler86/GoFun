package main

import "Blog/routers"

func main() {

	routers.InitRouter()

	//r :=  gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "success...",
	//	})
	//})
	//r.Run(":8080")
}