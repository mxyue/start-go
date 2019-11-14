package routes

import (
	"project/src/config"
	"project/src/middleware"
	devApi "project/src/routes/api/dev"
	v1Api "project/src/routes/api/v1"
	"syscall"

	"github.com/gin-gonic/gin"
)

//Router 路由入口
func Router() *gin.Engine {
	isNotProdEnv := !config.EnvIsProd()
	r := gin.New()

	if isNotProdEnv {
		setSwaggerRoute(r)
	}

	r.GET("/health/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": config.Version,
			"pid":     syscall.Getpid(),
		})
	})
	//logger放在下面,上面的不会打印出日志
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cross)

	v1Api.APIRouter(r)

	if isNotProdEnv {
		devApi.APIRouter(r)
	}

	return r
}
