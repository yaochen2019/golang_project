package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yaochen2019/gogofly/docs"
	"github.com/yaochen2019/gogofly/global"
)

// 对所有路由进行收集和注册
type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRegistRoute
)

// 注册路由
func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

// gin框架初始化
func InitRouter() {
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")
	InitBasePlatformRoutes()

	for _, fnRegistRoute := range gfnRouters {
		fnRegistRoute(rgPublic, rgAuth)
	}
	// 继承swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	stPort := viper.GetString("server.port")
	if stPort == "" {
		fmt.Sprint("失败")
		stPort = "8999"
	}
	err := r.Run()
	fmt.Sprint(":%s", stPort)
	if err != nil {
		global.Logger.Error(fmt.Sprint("start server Error: %s", err.Error()))

	}
	global.Logger.Info("")
}

// 初始化路由
func InitBasePlatformRoutes() {
	InitUserRoutes()
}
