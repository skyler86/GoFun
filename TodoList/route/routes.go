package route

import (
	"TodoList/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func NewRouter() *gin.Engine {
	gin.ForceConsoleColor()		// 此方法可以在控制台显示不同颜色的log日志信息

	r := gin.Default()		//创建gin的实例对象
	store := cookie.NewStore([]byte("something-very-secret"))		// 写入cookie
	r.Use(sessions.Sessions("mysession", store))		// 用session将store进行存储
	v1 := r.Group("api/v1") 			// 基础路由
	{
		// 用户的操作,定义两个接口
		v1.POST("user/register", controller.UserRegister)	// 注册接口
		v1.POST("user/login",controller.UserLogin)			// 登录接口
	}
	return r
}