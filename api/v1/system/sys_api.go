package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	"github/shansec/go-vue-admin/utils"
	SystemVerify "github/shansec/go-vue-admin/verify/system"
)

type SysApi struct{}

// CreateApi
// @Summary 创建 api
// @Tags SysApi
// @Accept json
// @Produce json
// @Param createApiInfo body system.SysApi true "创建 api"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/createApi [POST]
func (s *SysApi) CreateApi(c *gin.Context) {
	var createApiInfo system.SysApi
	_ = c.ShouldBindJSON(&createApiInfo)

	// 验证数据是否合法
	if err := utils.Verify(createApiInfo, SystemVerify.CreateApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := apiService.CreateApiInfo(&createApiInfo); err != nil {
		global.MAY_LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApi
// @Summary 删除 api
// @Tags SysApi
// @Accept json
// @Produce json
// @Param deleteApiInfo body system.SysApi true "删除 api"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/deleteApi [DELETE]
func (s *SysApi) DeleteApi(c *gin.Context) {
	var deleteApiInfo system.SysApi
	_ = c.ShouldBindJSON(&deleteApiInfo)

	// 验证数据是否合法
	if err := utils.Verify(deleteApiInfo, SystemVerify.DeleteApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := apiService.DeleteApiInfo(&deleteApiInfo); err != nil {
		global.MAY_LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateApi
// @Summary 更新 api
// @Tags SysApi
// @Accept json
// @Produce json
// @Param updateApiInfo body system.SysApi true "更新 api"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /api/updateApi [PUT]
func (s *SysApi) UpdateApi(c *gin.Context) {
	var updateApiInfo system.SysApi
	_ = c.ShouldBindJSON(&updateApiInfo)

	// 验证数据是否合法
	if err := utils.Verify(updateApiInfo, SystemVerify.CreateApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := apiService.UpdateApiInfo(&updateApiInfo); err != nil {
		global.MAY_LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetApiList
// @Summary 获取 api 列表
// @Description 分页查询 api 信息列表
// @Tags SysApi
// @Accept json
// @Produce json
// @Param  apiPageInfo body systemReq.GetApiList true "api 列表查询参数"
// @Success 200 {object} response.PageResult{list=[]system.SysApi, msg=string}	"api 列表获取成功"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "获取部门列表失败"
// @Router /api/getApiList [POST]
func (s *SysApi) GetApiList(c *gin.Context) {
	var apiPageInfo systemReq.GetApiList
	err := c.ShouldBindJSON(&apiPageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询 api 列表
	if apis, total, err := apiService.GetApisInfo(apiPageInfo); err != nil {
		global.MAY_LOGGER.Error("api 列表获取失败", zap.Error(err))
		response.FailWithMessage("api 列表获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     apis,
			Total:    total,
			Page:     apiPageInfo.Page,
			PageSize: apiPageInfo.PagSize,
		}, "api 列表获取成功", c)
	}
}
