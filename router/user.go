package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaochen2019/gogofly/api"
	"net/http"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"msg": "Login Middle ware",
				})
			}
		}()) //为当前user路由定义一个中间件

		rgPublicUser.POST("/login", userApi.Login)
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{"data": []map[string]any{
				{"id": 1, "name": "zs"},
				{"id": 2, "name": "ls"},
			}})
		})
		rgAuthUser.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{"id": 1, "name": "zs"})
		})

	})

}
