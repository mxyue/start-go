package dev

import (
	"project/src/middleware"
	"github.com/gin-gonic/gin"
)

//APIRouter dev API
func APIRouter(r *gin.Engine) {
	api := r.Group("api/dev")
	api.Use(middleware.Response)


}
