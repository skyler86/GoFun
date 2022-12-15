package middleware

import (
	"TodoList/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// 实现JWT中间件，来验证token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
//		var data interface{}	// 如果要使用data变量需要在pkg/e包里定义，暂时不用
		token := c.GetHeader("Authorization")	// 将鉴权放在Authorization里
		if token == "" {	// 如果传过来的参数不对,也就是没有token，就返回404
			code = 404
		}else{
			claim,err := utils.ParseToken(token)	// 如果不为空在把它的token进行解析
			if err != nil {
				code = 403		// 返回403表示token是否无权限的，是假的
			}else if time.Now().Unix() > claim.ExpiresAt {	// 如果你现在的时间大于你设定的过期时间，则说明你的token已经过期了。
				code = 401
			}
		}
		if code != 200 {
			c.JSON(200,gin.H{
				"status":code,
				"msg":"Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}