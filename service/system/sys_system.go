package system

import (
	"go.uber.org/zap"

	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/utils"
)

type SystemConfigService struct{}

// GetServerInfo
// @author: [Shansec](https://github.com/shansec)
// @function: GetServerInfo
// @description: 获取服务器信息
// @param: nil
// @return: server *utils.Server, err error
func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.MAY_LOGGER.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.MAY_LOGGER.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.MAY_LOGGER.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
