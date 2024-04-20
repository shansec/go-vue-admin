package system

import (
	"errors"
	"fmt"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/template/auto_template"
	"github/shansec/go-vue-admin/utils"
	"github/shansec/go-vue-admin/utils/ast"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"text/template"
)

type AutoCodeService struct{}

type astInjectionMeta struct {
	path         string
	importCodeF  string
	structNameF  string
	packageNameF string
	groupName    string
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string
}

type autoPackage struct {
	path  string
	cache string
	name  string
}

const (
	autoPath           = "autocode_template/"
	autocodePath       = "resource/autocode_template"
	plugServerPath     = "resource/plug_template/server"
	plugWebPath        = "resource/plug_template/web"
	packageService     = "service/%s/enter.go"
	packageServiceName = "service"
	packageRouter      = "router/%s/enter.go"
	packageRouterName  = "router"
	packageAPI         = "api/v1/%s/enter.go"
	packageAPIName     = "api/v1"
)

var (
	packageInjectionMap map[string]astInjectionMeta
	injectionPaths      []injectionMeta
)

func Init(packageName string) {
	injectionPaths = []injectionMeta{
		{
			path:        filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SApi, packageName), "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path:        filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SRouter, packageName), "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path:        filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SService, packageName), "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}
	packageInjectionMap = map[string]astInjectionMeta{
		packageServiceName: {
			path:         filepath.Join(global.MAY_CONFIG.AutoCode.Root, "service", "enter.go"),
			importCodeF:  "github/shansec/go-vue-admin/%s/%s",
			packageNameF: "%s",
			groupName:    "ServiceGroup",
			structNameF:  "%sServiceGroup",
		},
		packageRouterName: {
			path:         filepath.Join(global.MAY_CONFIG.AutoCode.Root, "router", "enter.go"),
			importCodeF:  "github/shansec/go-vue-admin/%s/%s",
			packageNameF: "%s",
			groupName:    "RouterGroup",
			structNameF:  "%sRouterGroup",
		},
		packageAPIName: {
			path:         filepath.Join(global.MAY_CONFIG.AutoCode.Root, "api/v1", "enter.go"),
			importCodeF:  "github/shansec/go-vue-admin/%s/%s",
			packageNameF: "%s",
			groupName:    "ApiGroup",
			structNameF:  "%sApiGroup",
		},
	}
}

// CreateAutoCode
// @author: [Shansec](https://github.com/shansec)
// @function: CreateAutoCode
// @description: 创建代码包
// @param: s *system.SysAutoCode
// @return: error
func (autoCodeService *AutoCodeService) CreateAutoCode(s *system.SysAutoCode) error {
	if s.PackageName == "system" || s.PackageName == "" {
		return errors.New("不能使用保留的包名")
	}
	if !errors.Is(global.MAY_DB.Where("package_name = ?", s.PackageName).First(&system.SysAutoCode{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的包名")
	}

	if err := autoCodeService.CreatePackageCache(s.PackageName); err != nil {
		return err
	}
	return global.MAY_DB.Create(&s).Error
}

func (autoCodeService *AutoCodeService) CreatePackageCache(packageName string) error {
	Init(packageName)
	pendingCache := []autoPackage{
		{
			path:  packageService,
			cache: string(auto_template.Service),
			name:  packageServiceName,
		},
		{
			path:  packageRouter,
			cache: string(auto_template.Router),
			name:  packageRouterName,
		},
		{
			path:  packageAPI,
			cache: string(auto_template.Api),
			name:  packageAPIName,
		},
	}

	webCache := []string{
		filepath.Join(global.MAY_CONFIG.AutoCode.WRoot, global.MAY_CONFIG.AutoCode.WWeb, global.MAY_CONFIG.AutoCode.WApi),
		filepath.Join(global.MAY_CONFIG.AutoCode.WRoot, global.MAY_CONFIG.AutoCode.WWeb, global.MAY_CONFIG.AutoCode.WForm),
	}

	for i, pend := range pendingCache {
		pendingCache[i].path = filepath.Join(global.MAY_CONFIG.AutoCode.Root, filepath.Clean(fmt.Sprintf(pend.path, packageName)))
	}

	for _, pend := range pendingCache {
		// 创建文件夹
		err := os.MkdirAll(filepath.Dir(pend.path), 0755)
		if err != nil {
			return err
		}
		file, err := os.Create(pend.path)
		if err != nil {
			return err
		}

		defer file.Close()

		parse, err := template.New("").Parse(pend.cache)
		if err != nil {
			return err
		}
		var packageStruct = struct {
			PackageName string `json:"package_name"`
		}{packageName}
		err = parse.Execute(file, packageStruct)
		if err != nil {
			return err
		}
	}

	// 插入结构代码
	for _, pend := range pendingCache {
		meta := packageInjectionMap[pend.name]
		if err := ast.ImportReference(meta.path, fmt.Sprintf(meta.importCodeF, pend.name, packageName), fmt.Sprintf(meta.structNameF, utils.FirstUpper(packageName)), fmt.Sprintf(meta.packageNameF, packageName), meta.groupName); err != nil {
			return err
		}
	}

	for _, web := range webCache {
		err := os.MkdirAll(filepath.Join(web, packageName), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
