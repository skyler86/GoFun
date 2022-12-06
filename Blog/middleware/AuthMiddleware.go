package middleware

import (
	"Blog/controllers"
	"Blog/types"
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

//AuthMiddleware
func AuthMiddleware() func(c *gin.Context)  {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.Error(c, int(types.ApiCode.NOAUTH), types.ApiCode.GetMessage(types.ApiCode.NOAUTH))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.Error(c, int(types.ApiCode.AUTHFORMATERROR), types.ApiCode.GetMessage(types.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}

		mc, err := controllers.ParseToken(parts[1])
		if err != nil {
			utils.Error(c, int(types.ApiCode.INVALIDTOKEN), types.ApiCode.GetMessage(types.ApiCode.INVALIDTOKEN))
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}