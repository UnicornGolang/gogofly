package api

import (
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

// api 定义
type UserApi struct{
  BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
    BaseApi: NewBaseApi(),
  }
}

// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录信息详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg": "Login success",
	// })
	// -----------------------------------------
	// OK(ctx, ResponseJson {
	//   Msg: "Login success",
	// })

	//==========================================================================
	// 参数绑定
	// ---------------
	// > ShouldBind() 将传递的 JSON 格式的数据与结构体绑定,不满足条件返回错误信息
	// > Bind() 将数据绑定到结构体，不满足绑定的条件则直接返回客户端请求

	// -------------------------------------------------------------------------
	// Should* 开头的方法会返回错误给对应的调用场景，让调用者自己处理抛出的异常
	// Bind* 开头的方法会直接 AbortWithError(), 状态码 400，直接返回 客户端

  //==========================================================================
  // 移动到 BaseAPI 中统一处理
  // -------------------------
	var userDTO dto.UserLoginDTO
	// if errs := ctx.ShouldBind(&userDTO); errs != nil {
	// 	Fail(ctx, ResponseJson{
	// 		// 返回更加友好的错误提示信息
	// 		// Msg: errs.Error(),
	// 		Msg: parseValidateError(errs.(validator.ValidationErrors), &userDTO).Error(),
	// 	})
	// 	return
	// }
  if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userDTO}).GetError(); err != nil {
    return
  }

	m.OK(ResponseJson{
		Data: userDTO,
	})
}

// @Tags 用户管理
// @Summary 用户注册
// @Description 用户注册信息详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "用户密码"
func (m UserApi) Register(ctx *gin.Context) {
	Fail(ctx, ResponseJson{
		Code: 9001,
		Msg:  "Register Error",
	})
}
