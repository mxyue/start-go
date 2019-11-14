package helper

import (
	"project/src/utils/codewrap"

	"github.com/gin-gonic/gin"
)

//ParamsError 设置参数错误信息
func ParamsError(c *gin.Context, message string) {
	c.Set("code", codewrap.Codes.ParamsError)
	c.Set("message", message)
}

//SystemError 系统错误
func SystemError(c *gin.Context, message string) {
	c.Set("code", codewrap.Codes.SystemError)
	c.Set("message", message)
}

//CustomError 自定义弹出消息错误
func CustomError(c *gin.Context, message string) {
	c.Set("code", codewrap.Codes.SystemError)
	c.Set("message", message)
}
