package api

import "github.com/gin-gonic/gin"

// api 定义
type AbortApi struct{}

// 构建一个测试 AbortWithStatusJSON() 函数的 API
func NewAbortApi() AbortApi {
	return AbortApi{}
}

// @Tags Abort
// @Description Abort 测试用例
// @Param name formData string "用户名称"
// @Param password formData string "用户密码"
// @Router /api/v1/public/abort/accept [post]
func (m AbortApi) Accept(ctx *gin.Context) {
	OK(ctx, ResponseJson{
		Msg: "Accept Api",
	})
}

func (m AbortApi) Abort(ctx *gin.Context) {
	OK(ctx, ResponseJson{
		Msg: "Abort Api",
	})
}
