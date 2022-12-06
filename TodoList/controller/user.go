package controller		// 可以理解成 api 层

import (
	"TodoList/logic"
	"github.com/gin-gonic/gin"
)

// 实现注册接口的函数
func UserRegister(c *gin.Context)  {	// 上下文对传过来的值进行绑定
	var userRegister logic.UserService		//声明一个user的服务对象变量
	if err := c.ShouldBind(&userRegister);err == nil {	// 绑定服务对象，将服务对象的值传过来
		res := userRegister.Register()		// 执行注册方法
		c.JSON(200,res)

	}else {
		c.JSON(400,err)
	}
}

// 实现登录接口的函数
func UserLogin(c *gin.Context)  {
	var userLogin logic.UserService		//声明一个user的服务对象变量
	if err := c.ShouldBind(&userLogin);err == nil {		// 绑定服务对象，将服务对象的值传过来
		res := userLogin.Login()		// 执行登录方法
		c.JSON(200,res)

	}else {
		c.JSON(400,err)
	}
}