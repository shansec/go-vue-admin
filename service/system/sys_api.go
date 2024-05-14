package system

import (
	"errors"
	"fmt"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	"gorm.io/gorm"
)

type ApiService struct{}

// CreateApiInfo
// @author: [Shansec](https://github.com/shansec)
// @function: CreateApiInfo
// @description: 创建 api
// @param: createApiInfo *system.SysApi
// @return: error
func (apiService *ApiService) CreateApiInfo(createApiInfo *system.SysApi) error {
	if !errors.Is(global.MAY_DB.Where("path = ?", createApiInfo.Path).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的路径")
	}
	return global.MAY_DB.Create(&createApiInfo).Error
}

// DeleteApiInfo
// @author: [Shansec](https://github.com/shansec)
// @function: DeleteApiInfo
// @description: 删除 api
// @param: deleteApiInfo *system.SysApi
// @return: error
func (apiService *ApiService) DeleteApiInfo(deleteApiInfo *system.SysApi) error {
	if err := global.MAY_DB.Delete(&deleteApiInfo).Error; err != nil {
		return errors.New("删除 api 信息失败")
	}
	return nil
}

// UpdateApiInfo
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateApiInfo
// @description: 更新 api
// @param: updateApiInfo *system.SysApi
// @return: error
func (apiService *ApiService) UpdateApiInfo(updateApiInfo *system.SysApi) error {
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
func (apiService *ApiService) GetApisInfo(getApisInfo systemReq.GetApiList) (apiList []system.SysApi, total int64, err error) {
	var apis []system.SysApi
	limit := getApisInfo.PagSize
	offset := getApisInfo.PagSize * (getApisInfo.Page - 1)
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
	err = db.Limit(limit).Offset(offset).Find(&apis).Error
	if err != nil {
		return nil, 0, errors.New("api 列表获取失败")
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, errors.New("api 列表获取失败")
	}
	return apis, total, nil
}
