// Package middlewares 中间件处理函数
package middlewares

import (
	"net/http"

	"github.com/yeongbok77/video-editor/controller"
	"github.com/yeongbok77/video-editor/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(http.StatusOK, controller.CodeMsgMap[controller.CodeNeedLogin])
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(token) // 解析token
		if err != nil {
			c.JSON(http.StatusOK, controller.CodeMsgMap[controller.CodeNeedLogin])
			c.Abort()
			return
		}
		// 后续的处理函数可以用过c.Get("user_id") 来获取当前请求的用户信息
		c.Set("user_id", mc.UserID)
		c.Next()
	}
}
