package api

import (
	global "gogofly/global/constants"
	"gogofly/service"
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

type HostApi struct {
	BaseApi
	Service *service.HostService
}

func NewHostApi() HostApi {
	return HostApi{
		Service: service.NewHostService(),
	}
}

// 关机
func (m HostApi) Shutdown(c *gin.Context) {
	var shutdownDTO dto.ShowdownDTO
	// 处理请求参数并校验
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &shutdownDTO}).GetError(); err != nil {
		return
	}
  // 处理业务逻辑并返回
  err := m.Service.Showdown(shutdownDTO)
  if err != nil {
    m.Fail(ResponseJson{
      Code: 10001,
      Msg: err.Error(),
    })
  }
  m.OK(ResponseJson{
    Msg: "shutdown success",
  })
  global.Log.Info("shutdown success")
}
