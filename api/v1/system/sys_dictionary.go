package system

import (
	"github.com/gin-gonic/gin"
	"github.com/shansec/go-vue-admin/dao/common/response"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
	"go.uber.org/zap"
)

type DictionaryApi struct{}

// CreateDictionary
// @Summary 添加字典
// @Description 添加字典，返回添加结果
// @Tags SysDictionary
// @Produce json
// @Param   dictionaryInfo body system.SysDictionary true "添加字典"
// @Success 200 {object} response.Response{msg=string}	"添加字典,返回添加结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "添加字典失败"
// @Router /dictionary/createDictionary [POST]
func (d *DictionaryApi) CreateDictionary(c *gin.Context) {
	var dictionaryInfo system.SysDictionary
	err := c.ShouldBindJSON(&dictionaryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.CreateDictionaryService(dictionaryInfo)
	if err != nil {
		global.MAY_LOGGER.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDictionary
// @Summary 删除字典
// @Description 删除字典，返回添加结果
// @Tags SysDictionary
// @Produce json
// @Param   dictionaryInfo body system.SysDictionary true "删除字典"
// @Success 200 {object} response.Response{msg=string}	"删除字典,返回删除结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "删除字典失败"
// @Router /dictionary/deleteDictionary [DELETE]
func (d *DictionaryApi) DeleteDictionary(c *gin.Context) {
	var dictionaryInfo system.SysDictionary
	err := c.ShouldBindJSON(&dictionaryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.DeleteDictionaryService(dictionaryInfo)
	if err != nil {
		global.MAY_LOGGER.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateDictionary
// @Summary 修改字典
// @Description 修改字典，返回修改结果
// @Tags SysDictionary
// @Produce json
// @Param   dictionaryInfo body system.SysDictionary true "修改字典"
// @Success 200 {object} response.Response{msg=string}	"修改字典,返回修改结果"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "修改字典失败"
// @Router /dictionary/updateDictionary [PUT]
func (d *DictionaryApi) UpdateDictionary(c *gin.Context) {
	var dictionaryInfo system.SysDictionary
	err := c.ShouldBindJSON(&dictionaryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.UpdateDictionaryService(&dictionaryInfo)
	if err != nil {
		global.MAY_LOGGER.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetDictionary
// @Summary 获取字典详情
// @Description 获取字典详情
// @Tags SysDictionary
// @Produce json
// @Param   dictionaryInfo body system.SysDictionary true "获取字典详情"
// @Success 200 {object} response.Response{data=system.SysDictionary, msg=string}	"获取字典详情"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "获取字典详情失败"
// @Router /dictionary/getDictionary [POST]
func (d *DictionaryApi) GetDictionary(c *gin.Context) {
	var dictionaryInfo system.SysDictionary
	err := c.ShouldBindJSON(&dictionaryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDictionary, err := dictionaryService.GetDictionaryService(dictionaryInfo.Type, dictionaryInfo.ID, dictionaryInfo.Status)
	if err != nil {
		global.MAY_LOGGER.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resDictionary": sysDictionary}, "查询成功", c)
}

// GetDictionaryInfoList
// @Summary 分页获取字典
// @Description 分页获取字典
// @Tags SysDictionary
// @Produce json
// @Param   dictionaryInfo body system.SysDictionary true "分页获取字典"
// @Success 200 {object} response.Response{data=system.SysDictionary, msg=string}	"分页获取字典"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response   "分页获取字典失败"
// @Router /dictionary/getDictionaryInfoList [POST]
func (d *DictionaryApi) GetDictionaryInfoList(c *gin.Context) {
	list, err := dictionaryService.GetDictionaryInfoListService()
	if err != nil {
		global.MAY_LOGGER.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}
