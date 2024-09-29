package system

import (
	"github.com/shansec/go-vue-admin/dao/request"
	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
)

type DictionaryDetailService struct{}

// CreateDictionaryDetailService
// @author: [Shansec](https://github.com/shansec)
// @function: CreateDictionaryDetailService
// @description: 创建字典详情
// @param: sysDictionaryDetail system.SysDictionaryDetail
// @return: err error
func (dictionaryDetailService *DictionaryDetailService) CreateDictionaryDetailService(sysDictionaryDetail system.SysDictionaryDetail) error {
	err := global.MAY_DB.Create(&sysDictionaryDetail).Error
	return err
}

// DeleteDictionaryDetailService
// @author: [Shansec](https://github.com/shansec)
// @function: DeleteDictionaryDetailService
// @description: 删除字典详情
// @param: sysDictionaryDetail system.SysDictionaryDetail
// @return: err error
func (dictionaryDetailService *DictionaryDetailService) DeleteDictionaryDetailService(sysDictionaryDetail system.SysDictionaryDetail) error {
	err := global.MAY_DB.Delete(&sysDictionaryDetail).Error
	return err
}

// UpdateDictionaryDetailService
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateDictionaryDetailService
// @description: 更新字典详情
// @param: sysDictionaryDetail system.SysDictionaryDetail
// @return: error
func (dictionaryDetailService *DictionaryDetailService) UpdateDictionaryDetailService(sysDictionaryDetail system.SysDictionaryDetail) error {
	err := global.MAY_DB.Save(sysDictionaryDetail).Error
	return err
}

// GetDictionaryDetailService
// @author: [Shansec](https://github.com/shansec)
// @function: GetDictionaryDetailService
// @description: 获取指定字典详情
// @param: id uint
// @return: sysDictionaryDetail system.SysDictionaryDetail, err error
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryDetailService(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.MAY_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

// GetDictionaryDetailListService
// @author: [Shansec](https://github.com/shansec)
// @function: GetDictionaryDetailListService
// @description: 分页获取字典详情
// @param: info request.SysDictionaryDetailSearch
// @return: list interface{}, total int64, err error
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryDetailListService(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MAY_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 有条件搜索
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}

func (dictionaryDetailService *DictionaryDetailService) GetDictionaryListService(dictionaryID uint) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	err = global.MAY_DB.Find(&sysDictionaryDetails, "sys_dictionary_id = ?", dictionaryID).Error
	return sysDictionaryDetails, err
}

func (dictionaryDetailService *DictionaryDetailService) GetDictionaryListByTypeService(t string) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.MAY_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.Find(&sysDictionaryDetails, "type = ?", t).Error
	return sysDictionaryDetails, err
}

func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByTypeValueService(t string, value uint) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.MAY_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.First(&sysDictionaryDetails, "type = ? AND value = ?", t, value).Error
	return sysDictionaryDetails, err
}
