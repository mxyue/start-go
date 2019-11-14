package v1

import (
	"project/src/middleware"

	"github.com/gin-gonic/gin"
)

//swagger中使用
type simpleResponse struct {
	//状态码
	code int
	//数据
	data interface{}
}

//APIRouter API v1 router
func APIRouter(r *gin.Engine) {

	api := r.Group("api/v1") //无校验路由
	api.Use(middleware.Response)

	authAPI := api.Group("") //校验的路由，未登录无法使用
	authAPI.Use(middleware.UserStrictAuth)

}
