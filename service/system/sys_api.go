package system

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	systemReq "github.com/shansec/go-vue-admin/dao/request"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
)

type ApiService struct{}

// CreateApiService
// @author: [Shansec](https://github.com/shansec)
// @function: CreateApiService
// @description: 创建 api
// @param: createApiInfo *system.SysApi
// @return: error
func (apiService *ApiService) CreateApiService(createApiInfo *system.SysApi) error {
	if !errors.Is(global.MAY_DB.Where("path = ?", createApiInfo.Path).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的路径")
	}
	return global.MAY_DB.Create(&createApiInfo).Error
}

// DeleteApiService
// @author: [Shansec](https://github.com/shansec)
// @function: DeleteApiService
// @description: 删除 api
// @param: deleteApiInfo *system.SysApi
// @return: error
func (apiService *ApiService) DeleteApiService(deleteApiInfo *system.SysApi) error {
	if err := global.MAY_DB.Delete(&deleteApiInfo).Error; err != nil {
		return errors.New("删除 api 信息失败")
	}
	return nil
}

// UpdateApiService
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateApiService
// @description: 更新 api
// @param: updateApiInfo *system.SysApi
// @return: error
func (apiService *ApiService) UpdateApiService(updateApiInfo *system.SysApi) error {
	var api system.SysApi
	err := global.MAY_DB.Model(&api).Where("id = ?", updateApiInfo.ID).Updates(map[string]interface{}{
		"path":        updateApiInfo.Path,
		"description": updateApiInfo.Description,
		"api_group":   updateApiInfo.ApiGroup,
		"method":      updateApiInfo.Method,
	}).Error
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return errors.New("更新 api 信息失败")
	}
	return nil
}

// GetApisInfo
// @author: [Shansec](https://github.com/shansec)
// @function: GetApisInfo
// @description: 获取 api 列表
// @param: getApisInfo systemReq.GetApiList
// @return: apiList []system.SysApi, total int64, err error
func (apiService *ApiService) GetApiListService(getApisInfo systemReq.GetApiList) (apiList []system.SysApi, total int64, err error) {
	var apis []system.SysApi
	limit := getApisInfo.PageSize
	offset := getApisInfo.PageSize * (getApisInfo.Page - 1)
	db := global.MAY_DB.Model(&system.SysApi{})
	if getApisInfo.Path != "" {
		db = db.Where("path LIKE ?", "%"+getApisInfo.Path+"%")
	}
	if getApisInfo.Description != "" {
		db = db.Where("description LIKE ?", "%"+getApisInfo.Description+"%")
	}
	if getApisInfo.ApiGroup != "" {
		db = db.Where("api_group = ?", getApisInfo.ApiGroup)
	}
	if getApisInfo.Method != "" {
		db = db.Where("method = ?", getApisInfo.Method)
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, errors.New("api 列表获取失败")
	}
	err = db.Limit(limit).Offset(offset).Find(&apis).Error
	if err != nil {
		return nil, 0, errors.New("api 列表获取失败")
	}
	return apis, total, nil
}
