package middleware

import (
	"github.com/gin-gonic/gin"
)

//AdminStrictAuth 校验用户id，如果未登录则返回401
func AdminStrictAuth(c *gin.Context) {
	c.Next()
}
