package system

import (
	"errors"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	systemReq "github/shansec/go-vue-admin/model/system/request"
	"gorm.io/gorm"
)

type DeptService struct{}

// EstablishDept
// @author: [Shansec](https://github.com/shansec)
// @function: EstablishDept
// @description: 添加部门
// @param: d system.SysDept
// @return: err error
func (deptService *DeptService) EstablishDept(d system.SysDept) (err error) {
	var dept system.SysDept
	if !errors.Is(global.MAY_DB.Where("dept_name = ?", d.DeptName).First(&dept).Error, gorm.ErrRecordNotFound) {
		return errors.New("部门名称已被占用")
	}
	err = global.MAY_DB.Create(&d).Error
	return err
}

// GetDept
// @author: [Shansec](https://github.com/shansec)
// @function: GetDept
// @description: 获取部门列表
// @param: info systemReq.GetDeptList
// @return: deptList []system.SysDept, total int64, err error
func (deptService *DeptService) GetDept(info systemReq.GetDeptList) (deptList []system.SysDept, total int64, err error) {
	var depts []system.SysDept
	limit := info.PagSize
	offset := info.PagSize * (info.Page - 1)
	db := global.MAY_DB.Model(&system.SysDept{})
	if info.DeptName != "" {
		db = db.Where("dept_name LIKE ?", "%"+info.DeptName+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Limit(limit).Offset(offset).Find(&depts).Error
	if err != nil {
		return nil, 0, errors.New("获取用户列表失败")
	}
	for _, dept := range depts {
		if dept.ParentId != 0 {
			continue
		}
		deptResult := deptService.GetDeptCall(depts, dept)
		deptList = append(deptList, deptResult)
	}
	return deptList, int64(len(deptList)), nil
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
