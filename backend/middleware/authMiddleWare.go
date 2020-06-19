package middleware

import (
	"blog/common"
	"blog/dto"
	"blog/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:] // 提取token的有效部分，因为 "Bearer "占了7位

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid { // 解析失败或者权限失效
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 验证通过获取claim中的userId
		userId := claims.UserId
		user := dto.GetUserById(userId)

		if user.Id == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 用户存在 将user信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}
