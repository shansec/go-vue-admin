package system

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
)

type DeptService struct{}

const DEPT_STATUS = "2"

// CreateDeptService
// @author: [Shansec](https://github.com/shansec)
// @function: CreateDeptService
// @description: 添加部门
// @param: d system.SysDept
// @return: err error
func (deptService *DeptService) CreateDeptService(d system.SysDept) (err error) {
	var dept system.SysDept
	if !errors.Is(global.MAY_DB.Where("dept_name = ?", d.DeptName).First(&dept).Error, gorm.ErrRecordNotFound) {
		return errors.New("部门名称已被占用")
	}
	err = global.MAY_DB.Create(&d).Error
	if err != nil {
		return err
	}
	deptPath := strconv.Itoa(d.DeptId) + "/"
	if d.ParentId != 0 {
		var deptParent system.SysDept
		global.MAY_DB.First(&deptParent, d.ParentId)
		deptPath = deptParent.DeptPath + deptPath
	} else {
		deptPath = "/0/" + deptPath
	}
	if err = global.MAY_DB.Model(&dept).Where("dept_id = ?", d.DeptId).Update("dept_path", deptPath).Error; err != nil {
		return err
	}
	return nil
}

// GetDeptListService
// @author: [Shansec](https://github.com/shansec)
// @function: GetDeptListService
// @description: 获取部门列表
// @param: info systemReq.GetDeptList
// @return: deptList []system.SysDept, total int64, err error
func (deptService *DeptService) GetDeptListService(info systemReq.GetDeptList) (deptList []system.SysDept, total int64, err error) {
	var depts []system.SysDept
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MAY_DB.Model(&system.SysDept{})
	if info.DeptName != "" {
		db = db.Where("dept_name LIKE ?", "%"+info.DeptName+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Limit(limit).Offset(offset).Order("sort ASC").Find(&depts).Error
	if err != nil {
		return nil, 0, errors.New("获取用户列表失败")
	}
	if info.DeptName == "" && info.Status == "" {
		for _, dept := range depts {
			if dept.ParentId != 0 {
				continue
			}
			deptResult := deptService.GetDeptCall(depts, dept)
			deptList = append(deptList, deptResult)
		}
		return deptList, int64(len(deptList)), nil
	} else {
		return depts, int64(len(depts)), nil
	}
}

// GetDeptCall 循环处理部门数据
func (deptService *DeptService) GetDeptCall(deptList []system.SysDept, dept system.SysDept) system.SysDept {
	var deptCalls []system.SysDept
	lists := deptList
	for _, list := range lists {
		if dept.DeptId != list.ParentId {
			continue
		}
		deptCall := system.SysDept{}
		deptCall.DeptId = list.DeptId
		deptCall.ParentId = list.ParentId
		deptCall.DeptPath = list.DeptPath
		deptCall.DeptName = list.DeptName
		deptCall.Sort = list.Sort
		deptCall.Leader = list.Leader
		deptCall.Phone = list.Phone
		deptCall.Email = list.Email
		deptCall.Status = list.Status
		deptCall.CreatedAt = list.CreatedAt
		dc := deptService.GetDeptCall(deptList, deptCall)
		deptCalls = append(deptCalls, dc)
	}
	dept.Children = deptCalls
	return dept
}

// DelDeptInfoService
// @author: [Shansec](https://github.com/shansec)
// @function: DelDeptInfoService
// @description: 删除部门信息
// @param: dept system.SysDept
// @return: err error
func (deptService *DeptService) DelDeptInfoService(dept system.SysDept) (err error) {
	var depts []system.SysDept
	var depart system.SysDept
	global.MAY_DB.Where("parent_id = ?", dept.DeptId).Find(&depts)
	if len(depts) != 0 {
		return errors.New("包含下级不能，请先删除下级部门")
	}
	if err = global.MAY_DB.Where("dept_id = ?", dept.DeptId).Delete(&depart).Error; err != nil {
		return errors.New("删除部门信息失败")
	}
	return nil
}

// UpdateDeptInfoService
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateDeptInfoService
// @description: 更改部门信息
// @param: deptInfo *system.SysDept
// @return: err error
func (deptService *DeptService) UpdateDeptInfoService(deptInfo *system.SysDept) error {
	var dept system.SysDept
	err := global.MAY_DB.Model(&dept).
		Select("updated_at", "dept_name", "sort", "leader", "phone", "email", "status").
		Where("dept_id = ?", deptInfo.DeptId).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"dept_name":  deptInfo.DeptName,
			"sort":       deptInfo.Sort,
			"leader":     deptInfo.Leader,
			"phone":      deptInfo.Phone,
			"email":      deptInfo.Email,
			"status":     deptInfo.Status,
		}).Error
	if err != nil {
		return errors.New("更新部门信息失败")
	} else {
		return nil
	}
}
