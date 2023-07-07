package router

import (
	"gogofly/api"

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
			publicUser.POST("/register", userApi.AddUser)
		}

		// 设置共有前缀,
		authUser := auth.Group("user")
		{
			// 设置需要认证的 API
			authUser.GET("/:id", userApi.GetUserById)
			authUser.POST("/list", userApi.GetUserList)
			authUser.PUT("/:id", userApi.UpdateUser)
			authUser.DELETE("/:id", userApi.DelUserById)
		}
	})
}
