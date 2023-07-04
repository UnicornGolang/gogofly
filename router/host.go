package router

import (
	"gogofly/api"

	"github.com/gin-gonic/gin"
)

// 改服务类主要用于服务的管理，例如服务健康度检测, 起停机等
// 一件停机操作 (这里实现主要是实现远程执行一个指令，关闭目标计算机)
func InitHostRoutes() {

	RegisterRoute(func(public *gin.RouterGroup, auth *gin.RouterGroup) {

		// 载入 User 的 api
		hostApi := api.NewHostApi()

		// 设置共有前缀,
		authHost := auth.Group("manager")
		{
			// 设置需要认证的 API
			authHost.POST("/shutdown", hostApi.Shutdown)
		}
	})
}
