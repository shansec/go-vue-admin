package system

import (
	"errors"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github/shansec/go-vue-admin/dao/request"
	"github/shansec/go-vue-admin/global"
)

type CasbinService struct{}

var CasbinServiceNew = new(CasbinService)

// UpdateCasbin
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateCasbin
// @description: 更新 casbin
// @param: RoleID uint, casbinInfos []request.CasbinInfo
// @return: error
func (casbinService *CasbinService) UpdateCasbin(RoleID uint, casbinInfos []request.CasbinInfo) error {
	roleId := strconv.Itoa(int(RoleID))
	casbinService.ClearCasbin(0, roleId)
	rules := [][]string{}
	// 权限去重
	deDuplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := roleId + v.Path + v.Method
		if _, ok := deDuplicateMap[key]; !ok {
			deDuplicateMap[key] = true
			rules = append(rules, []string{roleId, v.Path, v.Method})
		}
	}

	enforcer := casbinService.Casbin()
	success, _ := enforcer.AddPolicies(rules)
	if !success {
		return errors.New("添加失败")
	}
	return nil
}

// UpdateCasbinApi
// @author: [Shansec](https://github.com/shansec)
// @function: UpdateCasbinApi
// @description: api 更新
// @param: oldPath, oldMethod, path, method string
// @return: error
func (casbinService *CasbinService) UpdateCasbinApi(oldPath, oldMethod, path, method string) error {
	err := global.MAY_DB.Model(&gormAdapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": path,
		"v2": method,
	}).Error
	enforcer := casbinService.Casbin()
	err = enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

// GetPolicyPathByRoleId
// @author: [Shansec](https://github.com/shansec)
// @function: GetPolicyPathByRoleId
// @description: 获取 casbin 列表
// @param: RoleId uint
// @return: pathMap []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByRoleId(RoleId uint) (pathMap []request.CasbinInfo) {
	enforcer := casbinService.Casbin()
	roleId := strconv.Itoa(int(RoleId))
	policies, _ := enforcer.GetFilteredPolicy(0, roleId)
	for _, policy := range policies {
		pathMap = append(pathMap, request.CasbinInfo{
			Path:   policy[1],
			Method: policy[2],
		})
	}
	return pathMap
}

// ClearCasbin
// @author: [Shansec](https://github.com/shansec)
// @function: ClearCasbin
// @description: 清除 casbin
// @param: v int, p ...string
// @return: bool
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	enforcer := casbinService.Casbin()
	result, err := enforcer.RemoveFilteredPolicy(v, p...)
	if err != nil {
		zap.L().Error("清除失败", zap.Error(err))
		return false
	}
	return result
}

// RemoveFilteredPolicy
// @author: [Shansec](https://github.com/shansec)
// @function: RemoveFilteredPolicy
// @description: 清除指定的 casbin
// @param: db *gorm.DB, roleId string
// @return: error
func (casbinService *CasbinService) RemoveFilteredPolicy(db *gorm.DB, roleId string) error {
	return db.Delete(&gormAdapter.CasbinRule{}, "v0 = ?", roleId).Error
}

func (casbinService *CasbinService) SyncPolicy(db *gorm.DB, roleId string, rule [][]string) error {
	err := casbinService.RemoveFilteredPolicy(db, roleId)
	if err != nil {
		return err
	}
	return casbinService.AddPolicy(db, rule)
}

func (casbinService *CasbinService) AddPolicy(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormAdapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormAdapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

func (CasbinService *CasbinService) FreshCasbin() (err error) {
	e := CasbinService.Casbin()
	err = e.LoadPolicy()
	return err
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormAdapter.NewAdapterByDB(global.MAY_DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
