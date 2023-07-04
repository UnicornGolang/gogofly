package service

import (
	"context"
	"fmt"
	global "gogofly/global/constants"
	"gogofly/service/dto"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
)

type HostService struct{}

var hostService *HostService

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

// github.com/apenella/go-ansible@v1.1.7
// 对应的需要执行对应指令的服务需要安装 `ansible` 命令
func (m *HostService) Showdown(shutdownDTO dto.ShowdownDTO) error {
	var errResult error
	host := shutdownDTO.HostIP
	global.Log.Info("接受到参数：" + host)

	// 关机业务操作
	// 构建连接操作主机的参数
	ansibleConnectionOption := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user.name"),
	}
	// 构建命令参数
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  fmt.Sprintf("%s,", host),
		ModuleName: "command",
		Args:       viper.GetString("ansible.shutdownHost.args"),
		ExtraVars: map[string]any{
			"ansible_password": viper.GetString("ansible.user.password"),
		},
	}
	// 构建命令
	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOption,
		StdoutCallback:    "oneline", // 执行命令的过程在命令行中输出
	}

	// 执行命令
	errResult = adhoc.Run(context.TODO())

	return errResult
}
