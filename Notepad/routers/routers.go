package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	// 接收用户的请求
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//定义api：v1版本
	v1Group := r.Group("v1")
	{
		// 增删改查的待办事项：
		// 1.添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 2.查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 3.修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 4.删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
