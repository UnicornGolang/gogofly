package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

// 判断返回对象是否为空
func (resp ResponseJson) isEmpty() bool {
	return reflect.DeepEqual(resp, ResponseJson{})
}

// 指定更加原生的返回函数, 兼容一些不需要传递
// responseJson 数据的场景
func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.isEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

// 成功的请求
func OK(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, http.StatusOK, resp)
}

// 失败的请求，
func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerFail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}

// 兼容自定义的返回状态码的场景
func buildStatus(resp ResponseJson, defaultStatus int) int {
	if resp.Status == 0 {
		return defaultStatus
	}
	return resp.Status
}
