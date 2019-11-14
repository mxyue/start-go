package middleware

import (
	"github.com/gin-gonic/gin"
)

//Cross 跨域处理中间件
func Cross(c *gin.Context) {
	origin := c.GetHeader("Origin")
	if origin == "null" {
		origin = "*"
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Type,authorization,cache-control,expires")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.JSON(200, gin.H{"cross": "allow"})
		return
	}
	c.Next()
}
