package api

import (
	"errors"
	"fmt"
	global "gogofly/global/constants"
	"gogofly/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 基础 API 定义，用于封装 API 层的通用代码,
type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
  return BaseApi {
    Logger: global.Log,
  }
}

// 用于接收用于封装请求的对象
type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

// 添加异常
func (m *BaseApi) AddError(err error) {
	m.Errors = utils.AppendError(m.Errors, err)
}

// 获取错误信息
func (m *BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	// 绑定请求上下文
	m.Ctx = option.Ctx
	// 绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = m.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = m.Ctx.ShouldBind(option.DTO)
		}
		if errResult != nil {
			errResult = m.parseValidateError(errResult, option.DTO)
			m.AddError(errResult)
			m.Fail(ResponseJson{
				Msg: m.GetError().Error(),
			})
		}
	}
	return m
}

// 分析元素对象上绑定的 tag 信息
func (m *BaseApi) parseValidateError(errs error , target any) error {
	var errResult error
  // 先判断错误的类型是否为校验字段异常的类型，如果不是直接返回
  validateErr, ok := errs.(validator.ValidationErrors)
  if !ok {
    return errs
  }
	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	// 遍历返回的校验错误信息
	for _, fieldErr := range validateErr {
		// 获取到存在校验错误的字段对应的定义(包含了tag信息)
		field, _ := fields.FieldByName(fieldErr.Field())
		// 构建自定义的错误提示字段 name
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		// 根据错误提示字段 name 获取到 tag 中自定义的错误信息
		errMessage := field.Tag.Get(errMessageTag)
		// 做错误信息的拼装 (取出错误信息,没有则取出默认的 message)
		// 对取出的错误信息包装成 error 追加到对应的错误列表中
		if len(errMessage) == 0 {
			errMessage = field.Tag.Get("message")
		}
		if len(errMessage) == 0 {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

// 对基础返回再次封装
func (m *BaseApi) OK(resp ResponseJson) {
  OK(m.Ctx, resp)
}

func (m *BaseApi) Fail(resp ResponseJson) {
  Fail(m.Ctx, resp)
}

func (m *BaseApi) ServerFail(resp ResponseJson) {
  ServerFail(m.Ctx, resp)
}


