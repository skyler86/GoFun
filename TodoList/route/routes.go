package route

import (
	"TodoList/controller"
	"TodoList/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func NewRouter() *gin.Engine {
	gin.ForceConsoleColor()		// 此方法可以在控制台显示不同颜色的log日志信息

	r := gin.Default()		//创建gin的实例对象（生成了一个WSGI应用程序实例）
	store := cookie.NewStore([]byte("something-very-secret"))		// 写入cookie
	r.Use(sessions.Sessions("mysession", store))		// 用session将store进行存储
	v1 := r.Group("api/v1") 			// 基础路由
	{
		// 用户的操作,定义两个接口
		v1.POST("user/register", controller.UserRegister)	// 注册接口
		v1.POST("user/login",controller.UserLogin)			// 登录接口

		// 对备忘录的操作（增删改查）
		// 用于登陆保护
		authed := v1.Group("/")		// 先新建一个分组，然后进行中间件的认证：JWT的鉴权
		// 当执行下面这些路由时就会先验证JWT中间件，看看到底有没有这个权限来进行访问或者请求，这个分组会先
		authed.Use(middleware.JWT())
		{
			authed.POST("task",controller.CreateTask)		// 创建备忘录的路由
			authed.GET("task/:id",controller.ShowTask)		// 展示备忘录的路由，传入前端的id
			authed.GET("tasks",controller.ListTask)			// 展示所有备忘录的路由
		}
	}
	return r
}