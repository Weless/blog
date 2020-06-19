package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin","*") // 允许前台的8080端口访问
		c.Writer.Header().Set("Access-Control-Max-age","86400") // 缓存时间
		c.Writer.Header().Set("Access-Control-Allow-Methods","*") // * 代表允许所有的方法
		c.Writer.Header().Set("Access-Control-Allow-Headers","*") // * 所有header
		c.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		if c.Request.Method == http.MethodOptions{ // 如果有option请求
			c.AbortWithStatus(200)
		}else{
			c.Next()
		}
	}
}