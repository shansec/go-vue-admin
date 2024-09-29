package system

import (
	"errors"

	"github.com/shansec/go-vue-admin/global"
	"github.com/shansec/go-vue-admin/model/system"
	"gorm.io/gorm"
)

type DictionaryService struct{}

// CreateDictionaryService
// @author: [Shansec](https://github.com/shansec)
// @function: CreateDictionaryService
// @description: 创建字典
// @param: sysDictionary system.SysDictionary
// @return: err error
func (dictionaryService *DictionaryService) CreateDictionaryService(sysDictionary system.SysDictionary) (err error) {
	if !errors.Is(global.MAY_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的 type，不允许创建")
	}
	err = global.MAY_DB.Create(&sysDictionary).Error
	return err
}

// DeleteDictionaryService
// @author: [Shansec](https://github.com/shansec)
// @function: DeleteDictionaryService
// @description: 删除字典
// @param: sysDictionary system.SysDictionary
// @return: err error
func (dictionaryService *DictionaryService) DeleteDictionaryService(sysDictionary system.SysDictionary) (err error) {
	err = global.MAY_DB.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("操作有误")
	}
	if err != nil {
		return err
	}
	err = global.MAY_DB.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}
	if sysDictionary.SysDictionaryDetails != nil {
		return global.MAY_DB.Where("sys_dictionary_id = ?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

// UpdateDictionaryService
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateDictionaryService
// @description: 修改字典
// @param: sysDictionary *system.SysDictionary
// @return: err error
func (dictionaryService *DictionaryService) UpdateDictionaryService(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	err = global.MAY_DB.Where("id = ?", sysDictionary.ID).First(&dict).Error
	if err != nil {
		global.MAY_LOGGER.Debug(err.Error())
		return errors.New("查询字段数据失败")
	}
	if dict.Type != sysDictionary.Type {
		if !errors.Is(global.MAY_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的 type，不允许创建")
		}
	}
	err = global.MAY_DB.Model(&dict).Updates(sysDictionaryMap).Error
	return err
}

// GetDictionaryService
// @author: [Shansec](https://github.com/shansec)
// @function: GetDictionaryService
// @description: 获取指定字典信息
// @param: Type string, Id uint, status *bool
// @return: sysDictionary system.SysDictionary, err error
func (dictionaryService *DictionaryService) GetDictionaryService(Type string, Id uint, status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = global.MAY_DB.Where("(type = ? OR id = ?) AND status = ?", Type, Id, flag).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true).Order("sort")
	}).First(&sysDictionary).Error
	return
}

// GetDictionaryInfoListService
// @author: [Shansec](https://github.com/shansec)
// @function: GetDictionaryInfoListService
// @description: 分页获取字典信息
// @param:
// @return: list interface{}, err error
func (dictionaryService *DictionaryService) GetDictionaryInfoListService() (list interface{}, err error) {
	var sysDictionaries []system.SysDictionary
	err = global.MAY_DB.Find(&sysDictionaries).Error
	return sysDictionaries, err
}
