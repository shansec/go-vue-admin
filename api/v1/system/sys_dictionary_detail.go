package system

import (
	"github.com/gin-gonic/gin"
	"github.com/shansec/go-vue-admin/dao/common/response"
	request "github.com/shansec/go-vue-admin/dao/request"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
	"github.com/shansec/go-vue-admin/utils"
	systemVerify "github.com/shansec/go-vue-admin/verify/system"
	"go.uber.org/zap"
)

type DictionaryDetailApi struct{}

// CreateDictionaryDetail
// @Summary 添加字典详情
// @Description 添加字典详情，返回添加结果
// @Tags SysDictionaryDetail
// @Produce json
// @Param   dictionaryDetailInfo body system.SysDictionaryDetail true "添加字典详情"
// @Success 200 {object} response.Response{msg=string}	"添加字典详情,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "添加字典详情失败"
// @Router /dictionaryDetail/createDictionaryDetail [POST]
func (d *DictionaryDetailApi) CreateDictionaryDetail(c *gin.Context) {
	var dictionaryDetailInfo system.SysDictionaryDetail
	err := c.ShouldBindJSON(&dictionaryDetailInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.CreateDictionaryDetailService(dictionaryDetailInfo)
	if err != nil {
		global.MAY_LOGGER.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDictionaryDetail
// @Summary 删除字典详情
// @Description 删除字典详情，返回删除结果
// @Tags SysDictionaryDetail
// @Produce json
// @Param   dictionaryDetailInfo body system.SysDictionaryDetail true "删除字典详情"
// @Success 200 {object} response.Response{msg=string}	"删除字典详情,返回删除结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "删除字典详情失败"
// @Router /dictionaryDetail/deleteDictionaryDetail [DELETE]
func (d *DictionaryDetailApi) DeleteDictionaryDetail(c *gin.Context) {
	var dictionaryDetailInfo system.SysDictionaryDetail
	err := c.ShouldBindJSON(&dictionaryDetailInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.DeleteDictionaryDetailService(dictionaryDetailInfo)
	if err != nil {
		global.MAY_LOGGER.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateDictionaryDetail
// @Summary 修改字典详情
// @Description 修改字典详情，返回修改结果
// @Tags SysDictionaryDetail
// @Produce json
// @Param   dictionaryDetailInfo body system.SysDictionaryDetail true "修改字典详情"
// @Success 200 {object} response.Response{msg=string}	"修改字典详情,返回修改结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "修改字典详情失败"
// @Router /dictionaryDetail/updateDictionaryDetail [PUT]
func (d *DictionaryDetailApi) UpdateDictionaryDetail(c *gin.Context) {
	var dictionaryDetailInfo system.SysDictionaryDetail
	err := c.ShouldBindJSON(&dictionaryDetailInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.UpdateDictionaryDetailService(dictionaryDetailInfo)
	if err != nil {
		global.MAY_LOGGER.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	response.OkWithMessage("更新成功", c)
}

// GetSysDictionaryDetail
// @Summary 获取字典详情
// @Description 获取字典详情
// @Tags SysDictionaryDetail
// @Produce json
// @Param   dictionaryDetailInfo body system.SysDictionaryDetail true "获取字典详情"
// @Success 200 {object} response.Response{data=system.SysDictionaryDetail, msg=string}	"获取字典详情"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "获取字典详情失败"
// @Router /dictionaryDetail/getDictionaryDetail [POST]
func (d *DictionaryDetailApi) GetSysDictionaryDetail(c *gin.Context) {
	var dictionaryDetailInfo system.SysDictionaryDetail
	err := c.ShouldBindJSON(&dictionaryDetailInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(dictionaryDetailInfo, systemVerify.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resDictionaryDetailInfo, err := dictionaryDetailService.GetDictionaryDetailService(dictionaryDetailInfo.ID)
	if err != nil {
		global.MAY_LOGGER.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	}
	response.OkWithDetailed(gin.H{"resDictionaryDetailInfo": resDictionaryDetailInfo}, "查询成功", c)
}

// GetDictionaryDetailList
// @Summary 分页获取字典详情
// @Description 分页获取字典详情
// @Tags SysDictionaryDetail
// @Produce json
// @Param   pageInfo body request.SysDictionaryDetailSearch true "分页获取字典详情"
// @Success 200 {object} response.Response{data=response.PageResult, msg=string}	"分页获取字典详情"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "分页获取字典详情失败"
// @Router /dictionaryDetail/getDictionaryDetailList [POST]
func (d *DictionaryDetailApi) GetDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dictionaryDetailService.GetDictionaryDetailListService(pageInfo)
	if err != nil {
		global.MAY_LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
