package router

import (
	"gogofly/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {

	RegisterRoute(func(public *gin.RouterGroup, auth *gin.RouterGroup) {

		// 载入 User 的 api
		userApi := api.NewUserApi()

		// 注册一个中间件,
		publicUser := public.Group("user") // 为了层次关系更加明确，使用括号来增加层次关系
		{
			// 设置开放 API
			publicUser.POST("/login", userApi.Login)
			publicUser.POST("/register", userApi.Register)
		}

		// 设置共有前缀,
		authUser := auth.Group("user")
		{
			// 设置需要认证的 API
			authUser.GET("", func(ctx *gin.Context) {
				// { data: { id: 1, name: "zs"}}
				ctx.JSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
						{"id": 2, "name": "ls"},
					},
				})
			})
			authUser.GET("/:id", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "zs",
				})
			})
		}
	})
}
