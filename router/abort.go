package router

import (
	"gogofly/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAbortRoutes() {

	RegisterRoute(func(public *gin.RouterGroup, auth *gin.RouterGroup) {

		// 载入 abort 的 api
		abortApi := api.NewAbortApi()

		// 注册一个中间件,
		abortPublic := public.Group("abort").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {
				//
				// AbortWithStatusJSON 与 JSON 之间的区别
				// --------------------------------------------------------------------------
				// JSON 之间的区别是，在使用一些中间件的时候能终止后续链路上的操作,
				// 当中间件阻断了请求，已经返回的情况下请求不会继续向下执行
				// > AbortWithStatusJSON : 当前中间件中已经返回了一个结果，
				//         则对应的请求不会继续往下执行  {"msg":"Login MiddleWare"}
				// > JSON : 即使在中间件中返回了一个结果, 程序会继续往下执行
				//         则返回值为 {"msg":"Login MiddleWare"}{"msg":"Login success"}
				// --------------------------------------------------------------------------
				// ctx.JSON(200, gin.H {
				ctx.AbortWithStatusJSON(200, gin.H {
					"msg": "Login MiddleWare",
				})
			}
		}())
		// 为了层次关系更加明确，使用括号来增加层次关系
		{
			// 设置开放 API
			abortPublic.POST("/accept", abortApi.Accept)
			abortPublic.POST("/abort", abortApi.Abort)
		}
		// 设置共有前缀,
		authAbort := auth.Group("abort")
		{
			// 设置需要认证的 API
			authAbort.GET("", func(ctx *gin.Context) {
				// { data: { id: 1, name: "zs"}}
				ctx.JSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
						{"id": 2, "name": "ls"},
					},
				})
			})
		}
	})
}
