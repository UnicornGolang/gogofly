package api

import (
	"fmt"
	"gogofly/service"
	"gogofly/service/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 按照模块将服务的错误码提取出来
const (
	ERR_CODE_ADD_USER      = 10011
	ERR_CODE_USRT_BY_ID    = 10012
	ERR_CODE_GET_USER_LIST = 10013
	ERR_CODE_UPDATE_UASER  = 10014
	ERR_CODE_DEL_UESR      = 10015
	ERR_CODE_LOGIN         = 10016
)

// api 定义
type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
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
	// -----------------------------------------

	// 移动到 BaseAPI 中统一处理
	// -----------------------------------------
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

	// 用户登录
	user, token, err := m.Service.Login(userDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ERR_CODE_LOGIN,
			Msg:    err.Error(),
		})
		return
	}
	if err != nil {
		m.Fail(ResponseJson{
			Code: ERR_CODE_LOGIN,
			Msg:  err.Error(),
		})
	}
	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}


// @Tags 用户管理
// @Summary 添加用户
// @Description 添加用户的详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "用户密码"
// @Param email formData string true "用户邮箱"
// @Param mobile formData string true "用户手机号码"
// @Param avatar formData file "用户头像"
func (m UserApi) AddUser(c *gin.Context) {
	var userDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &userDTO}).GetError(); err != nil {
		return
	}
	// 添加用户可能需要上传图片
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)
	filePath := fmt.Sprintf("./upload/%s", file.Filename)
	_ = c.SaveUploadedFile(file, filePath)
	// 图片的路径保存到数据库
	userDTO.Avatar = filePath
	err := m.Service.AddUser(&userDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: userDTO,
	})
}

// @Tags 用户管理
// @Summary 根据用户 Id 查找
// @Description 根据用户 Id 精确查找某个用户的详细信息
func (m UserApi) GetUserById(c *gin.Context) {
	var commonDTO dto.CommonDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &commonDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	//
	user, err := m.Service.GerUserById(&commonDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_USRT_BY_ID,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: user,
	})
}

// @Tags 用户管理
// @Summary 查找用户列表
// @Description 根据分页参数获取用户列表信息
func (m UserApi) GetUserList(c *gin.Context) {
	var userListDTO dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &userListDTO}).GetError(); err != nil {
		return
	}
	userList, total, err := m.Service.GetUserList(&userListDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data:  userList,
		Total: total,
	})
}

// @Tags 用户管理
// @Summary 用户信息更新
// @Description 用户信息更新，需要传入用户的 ID,更新用户的名称,电话,邮箱,头像等
// @Param id uri uint true "用户主键ID"
// @Param name json string false "用户名"
// @Param email json string true "用户邮箱"
// @Param mobile json string true "用户手机号码"
// @Param avatar json file "用户头像"
func (m UserApi) UpdateUser(c *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &userUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}
	err := m.Service.UpdateUser(userUpdateDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_UASER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{})
}

// @Tags 用户管理
// @Summary 根据 Id 删除用户
// @Summary 根据 Id 主键删除用户
func (m UserApi) DelUserById(c *gin.Context) {
	var commonDTO dto.CommonDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &commonDTO, BindUri: true}).GetError(); err != nil {
		return
	}
	err := m.Service.DelUserById(&commonDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_DEL_UESR,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{})
}
