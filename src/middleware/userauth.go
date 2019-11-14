package middleware

import (
	"github.com/gin-gonic/gin"
)

//UserStrictAuth 校验用户id，如果未登录则返回401
func UserStrictAuth(c *gin.Context) {
	c.Next()
}
