//+build noswagger

package routes

import (
	"github.com/gin-gonic/gin"
)

func setSwaggerRoute(r *gin.Engine) {
	r.GET("/swagger/*any", func(ctx *gin.Context) {
		ctx.String(404, "Not Fount")
	})
}
