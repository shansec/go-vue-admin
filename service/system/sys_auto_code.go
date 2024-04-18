package system

import (
	"errors"
	"fmt"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/utils"
	"gorm.io/gorm"
	"path/filepath"
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
			path: filepath.Join(global.MAY_CONFIG.AutoCode.Root,
				global.MAY_CONFIG.AutoCode.SServer, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SApi, packageName), "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path: filepath.Join(global.MAY_CONFIG.AutoCode.Root,
				global.MAY_CONFIG.AutoCode.SServer, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SRouter, packageName), "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path: filepath.Join(global.MAY_CONFIG.AutoCode.Root,
				global.MAY_CONFIG.AutoCode.SServer, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SService, packageName), "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}
	packageInjectionMap = map[string]astInjectionMeta{
		packageServiceName: {
			path:         filepath.Join(global.MAY_CONFIG.AutoCode.Root, global.MAY_CONFIG.AutoCode.SServer, "service", "enter,go"),
			importCodeF:  "github/shansec/go-vue-admin/%s/%s",
			packageNameF: "%s",
			groupName:    "ServiceGroup",
			structNameF:  "%sServiceGroup",
		},
		packageRouterName: {
			path:         filepath.Join(global.MAY_CONFIG.AutoCode.Root, global.MAY_CONFIG.AutoCode.SServer, "router", "enter,go"),
			importCodeF:  "github/shansec/go-vue-admin/%s/%s",
			packageNameF: "%s",
			groupName:    "RouterGroup",
			structNameF:  "%sRouterGroup",
		},
		packageAPIName: {
			path:         filepath.Join(global.MAY_CONFIG.AutoCode.Root, global.MAY_CONFIG.AutoCode.SServer, "api/v1", "enter,go"),
			importCodeF:  "github/shansec/go-vue-admin/%s/%s",
			packageNameF: "%s",
			groupName:    "ApiGroup",
			structNameF:  "%sApiGroup",
		},
	}
}

func (autoCodeService *AutoCodeService) CreateAutoCode(s *system.SysAutoCode) error {
	if s.PackageName == "system" || s.PackageName == "" {
		return errors.New("不能使用保留的包名")
	}
	if !errors.Is(global.MAY_DB.Where("package_name = ?", s.PackageName).First(&system.SysAutoCode{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的包名")
	}

}

func (autoCodeService *AutoCodeService) CreatePackageCache(packageName string) error {
	Init(packageName)
	pendingCache := []autoPackage{
		{
			path:  packageService,
			cache: string(utils.Service),
			name:  packageServiceName,
		},
		{
			path:  packageRouter,
			cache: string(utils.Router),
			name:  packageRouterName,
		},
		{
			path:  packageAPI,
			cache: string(utils.Api),
			name:  packageAPIName,
		},
	}

}
