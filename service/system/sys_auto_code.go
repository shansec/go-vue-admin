package system

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"gorm.io/gorm"

	systemReq "github/shansec/go-vue-admin/dao/request"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/system"
	"github/shansec/go-vue-admin/template/auto_template"
	"github/shansec/go-vue-admin/utils"
	"github/shansec/go-vue-admin/utils/ast"
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

type tplData struct {
	template         *template.Template
	autoPackage      string
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
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

// CreatePackageService
// @author: [Shansec](https://github.com/shansec)
// @function: CreatePackageService
// @description: 创建代码包
// @param: s *system.SysAutoCode
// @return: error
func (autoCodeService *AutoCodeService) CreatePackageService(s *system.SysAutoCode) error {
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

// GetPackageListService
// @author: [Shansec](https://github.com/shansec)
// @function: GetPackageListService
// @description: 获取代码包列表
// @param: s *system.SysAutoCode
// @return: error
func (autoCodeService *AutoCodeService) GetPackageListService(info systemReq.GetPackageList) (pkgList []system.SysAutoCode, total int64, err error) {
	var autoCodes []system.SysAutoCode
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MAY_DB.Model(&system.SysAutoCode{})

	if info.PackageName != "" {
		db = db.Where("package_name LIKE ?", "%"+info.PackageName+"%")
	}
	err = db.Limit(limit).Offset(offset).Find(&autoCodes).Error
	err = global.MAY_DB.Find(&pkgList).Error
	if err != nil {
		return nil, 0, errors.New("获取包列表失败")
	}
	return autoCodes, int64(len(autoCodes)), nil
}

// DelPackageService
// @author: [Shansec](https://github.com/shansec)
// @function: DelPackageService
// @description: 删除代码包
// @param: s *system.SysAutoCode
// @return: error
func (autoCodeService *AutoCodeService) DelPackageService(s *system.SysAutoCode) error {
	return global.MAY_DB.Delete(s).Error
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

// PreviewCodeService
// @author: [Shansec](https://github.com/shansec)
// @function: PreviewCodeService
// @description: 预览代码
// @param: a system.AutoCodeStruct
// @return: map[string]string, error
func (autoCodeService *AutoCodeService) PreviewCodeService(a system.AutoCodeStruct) (map[string]string, error) {
	makeDictType(&a)
	for i := range a.Fields {
		if a.Fields[i].FieldType == "time.Time" {
			a.HasTimer = true
			if a.Fields[i].FieldSearchType != "" {
				a.HasSearchTimer = true
			}
		}
		if a.Fields[i].Sort {
			a.NeedSort = true
		}
		if a.Fields[i].FieldType == "picture" {
			a.HasPic = true
		}
		if a.Fields[i].FieldType == "video" {
			a.HasPic = true
		}
		if a.Fields[i].FieldType == "richtext" {
			a.HasRichText = true
		}
		if a.Fields[i].FieldType == "pictures" {
			a.HasPic = true
			a.NeedJSON = true
		}
		if a.Fields[i].FieldType == "file" {
			a.HasFile = true
			a.NeedJSON = true
		}

		if a.DefaultModel {
			a.PrimaryField = &system.Field{
				FieldName:    "ID",
				FieldType:    "uint",
				FieldDesc:    "ID",
				FieldJson:    "ID",
				DataTypeLong: "20",
				Comment:      "主键ID",
				ColumnName:   "id",
			}
		}

		if !a.DefaultModel && a.PrimaryField == nil && a.Fields[i].PrimaryKey {
			// 设置主键
			a.PrimaryField = a.Fields[i]
		}
	}
	dataList, _, needMkdir, err := autoCodeService.getNeedList(&a)
	if err != nil {
		return nil, err
	}

	// 创建文件夹
	if err = utils.CreateDir(needMkdir...); err != nil {
		return nil, err
	}

	resultMap := make(map[string]string)

	for _, data := range dataList {
		ext := ""
		if ext = filepath.Ext(data.autoCodePath); ext != ".txt" {
			continue
		}

		file, err := os.OpenFile(data.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return nil, err
		}
		err = data.template.Execute(file, a)
		if err != nil {
			return nil, err
		}
		_ = file.Close()

		file, err = os.OpenFile(data.autoCodePath, os.O_CREATE|os.O_RDONLY, 0o755)
		if err != nil {
			return nil, err
		}
		builder := strings.Builder{}
		builder.WriteString("```")

		if ext != "" && strings.Contains(ext, ".") {
			builder.WriteString(strings.Replace(ext, ".", "", -1))
		}
		builder.WriteString("\n\n")
		readData, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		builder.WriteString(string(readData))
		builder.WriteString("\n\n````")
		paths := strings.Split(data.autoCodePath, string(os.PathSeparator))
		resultMap[paths[1]+"-"+paths[3]] = builder.String()
		_ = file.Close()
	}
	defer func() {
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	return resultMap, nil
}

func (autoCodeService *AutoCodeService) CreateCode(a system.AutoCodeStruct, menuID uint, ids ...uint) (err error) {
	makeDictType(&a)
	for i := range a.Fields {
		if a.Fields[i].FieldType == "time.Time" {
			a.HasTimer = true
			if a.Fields[i].FieldSearchType != "" {
				a.HasSearchTimer = true
			}
		}
		if a.Fields[i].Sort {
			a.NeedSort = true
		}
		if a.Fields[i].FieldType == "picture" {
			a.HasPic = true
		}
		if a.Fields[i].FieldType == "video" {
			a.HasPic = true
		}
		if a.Fields[i].FieldType == "richtext" {
			a.HasRichText = true
		}
		if a.Fields[i].FieldType == "pictures" {
			a.NeedJSON = true
			a.HasPic = true
		}
		if a.Fields[i].FieldType == "file" {
			a.NeedJSON = true
			a.HasFile = true
		}
		if a.DefaultModel {
			a.PrimaryField = &system.Field{
				FieldName:    "ID",
				FieldType:    "uint",
				FieldDesc:    "ID",
				FieldJson:    "ID",
				DataTypeLong: "20",
				Comment:      "主键ID",
				ColumnName:   "id",
			}
		}
		if !a.DefaultModel && a.PrimaryField == nil && a.Fields[i].PrimaryKey {
			a.PrimaryField = a.Fields[i]
		}
	}
	// 增加判断: 重复创建struct
	// if a.AutoMoveFile && AutoCodeHistoryServiceApp.Repeat(autoCode.BusinessDB, autoCode.StructName, autoCode.Package) {
	// 	return RepeatErr
	// }
	dataList, fileList, needMkdir, err := autoCodeService.getNeedList(&a)
	if err != nil {
		return err
	}
	// meta, _ := json.Marshal(a)

	// 增加判断：Package不为空
	if a.Package == "" {
		return errors.New("Package为空\n")
	}

	// 写入文件前，先创建文件夹
	if err = utils.CreateDir(needMkdir...); err != nil {
		return err
	}

	// 生成文件
	for _, value := range dataList {
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, a); err != nil {
			return err
		}
		_ = f.Close()
	}

	defer func() { // 移除中间文件
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	bf := strings.Builder{}
	idBf := strings.Builder{}
	injectionCodeMeta := strings.Builder{}
	for _, id := range ids {
		idBf.WriteString(strconv.Itoa(int(id)))
		idBf.WriteString(";")
	}
	if a.AutoMoveFile { // 判断是否需要自动转移
		Init(a.Package)
		for index := range dataList {
			autoCodeService.addAutoMoveFile(&dataList[index])
		}
		// 判断目标文件是否都可以移动
		for _, value := range dataList {
			if utils.FileExist(value.autoMoveFilePath) {
				return errors.New(fmt.Sprintf("目标文件已存在:%s\n", value.autoMoveFilePath))
			}
		}
		for _, value := range dataList { // 移动文件
			if err := utils.FileMove(value.autoCodePath, value.autoMoveFilePath); err != nil {
				return err
			}
		}

		{
			// 在gorm.go 注入 自动迁移
			path := filepath.Join(global.MAY_CONFIG.AutoCode.Root, global.MAY_CONFIG.AutoCode.SInitialize, "gorm.go")
			varDB := utils.MaheHump(a.BusinessDB)
			ast.AddRegisterTablesAst(path, "RegisterTables", a.Package, varDB, a.BusinessDB, a.StructName)
		}

		{
			// router.go 注入 自动迁移
			path := filepath.Join(global.MAY_CONFIG.AutoCode.Root, global.MAY_CONFIG.AutoCode.SInitialize, "router.go")
			ast.AddRouterCode(path, "Routers", a.Package, a.StructName)
		}
		// 给各个enter进行注入
		err = injectionCode(a.StructName, &injectionCodeMeta)
		if err != nil {
			return
		}
		// 保存生成信息
		for _, data := range dataList {
			if len(data.autoMoveFilePath) != 0 {
				bf.WriteString(data.autoMoveFilePath)
				bf.WriteString(";")
			}
		}
	} else { // 打包
		if err = utils.ZipFiles("./govueadmin.zip", fileList, ".", "."); err != nil {
			return err
		}
	}
	// if a.AutoMoveFile || a.AutoCreateApiToSql || a.AutoCreateMenuToSql {
	// 	if a.TableName != "" {
	// 		err = AutoCodeHistoryServiceApp.CreateAutoCodeHistory(
	// 			string(meta),
	// 			a.StructName,
	// 			a.Description,
	// 			bf.String(),
	// 			injectionCodeMeta.String(),
	// 			a.TableName,
	// 			idBf.String(),
	// 			a.Package,
	// 			a.BusinessDB,
	// 			menuID,
	// 		)
	// 	} else {
	// 		err = AutoCodeHistoryServiceApp.CreateAutoCodeHistory(
	// 			string(meta),
	// 			a.StructName,
	// 			a.Description,
	// 			bf.String(),
	// 			injectionCodeMeta.String(),
	// 			a.StructName,
	// 			idBf.String(),
	// 			a.Package,
	// 			a.BusinessDB,
	// 			menuID,
	// 		)
	// 	}
	// }
	if err != nil {
		return err
	}
	if a.AutoMoveFile {
		return errors.New("创建代码成功并移动文件成功")
	}
	return nil
}

func makeDictType(autoCode *system.AutoCodeStruct) {
	DictTypesM := make(map[string]string)
	for _, code := range autoCode.Fields {
		if code.DictType != "" {
			DictTypesM[code.DictType] = ""
		}
	}

	for key := range DictTypesM {
		autoCode.DictTypes = append(autoCode.DictTypes, key)
	}
}

func (autoCodeService *AutoCodeService) addAutoMoveFile(data *tplData) {
	base := filepath.Base(data.autoCodePath)
	fileSlice := strings.Split(data.autoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	// autocode_template/server/admin/api/admin.go 	n = 5
	// autocode_template/web/admin/api/admin.js
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SRouter, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SApi, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SService, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SModel, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root, fmt.Sprintf(global.MAY_CONFIG.AutoCode.SRequest, data.autoPackage), base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root,
				global.MAY_CONFIG.AutoCode.WWeb, global.MAY_CONFIG.AutoCode.WApi, data.autoPackage, base)
		} else if strings.Contains(fileSlice[n-2], "form") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root,
				global.MAY_CONFIG.AutoCode.WWeb, global.MAY_CONFIG.AutoCode.WForm, data.autoPackage, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			data.autoMoveFilePath = filepath.Join(global.MAY_CONFIG.AutoCode.Root,
				global.MAY_CONFIG.AutoCode.WWeb, global.MAY_CONFIG.AutoCode.WTable, data.autoPackage, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), base)
		}
	}
}

func (autoCodeService *AutoCodeService) CreateApiAuto(a *system.AutoCodeStruct) (ids []uint, err error) {
	apiList := []system.SysApi{
		{
			Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
			Description: "新增" + a.Description,
			ApiGroup:    a.Description,
			Method:      "POST",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
			Description: "删除" + a.Description,
			ApiGroup:    a.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
			Description: "批量删除" + a.Description,
			ApiGroup:    a.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
			Description: "更新" + a.Description,
			ApiGroup:    a.Description,
			Method:      "PUT",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
			Description: "根据ID获取" + a.Description,
			ApiGroup:    a.Description,
			Method:      "GET",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName,
			Description: "获取" + a.Description + "列表",
			ApiGroup:    a.Description,
			Method:      "GET",
		},
	}
	global.MAY_DB.Transaction(func(ctx *gorm.DB) error {
		for _, data := range apiList {
			var api system.SysApi
			if errors.Is(ctx.Where("path = ? AND method = ?", data.Path, data.Method).First(&api).Error, gorm.ErrRecordNotFound) {
				if err = ctx.Create(&data).Error; err != nil {
					return err
				} else {
					ids = append(ids, data.ID)
				}
			}
		}
		return nil
	})
	return ids, err
}

// func (autoCodeService *AutoCodeService) AutoCreateMenu(a *system.AutoCodeStruct) (id uint, err error) {
// 	var menu system.SysBaseMenu
// 	err = global.MAY_DB.First(&menu, "name = ?", a.Abbreviation).Error
// 	if err == nil {
// 		return 0, errors.New("存在相同的菜单路由，请关闭自动创建菜单功能")
// 	}
// 	menu.ParentId = 0
// 	menu.Name = a.Abbreviation
// 	menu.Path = a.Abbreviation
// 	menu.Meta.Title = a.Description
// 	menu.Component = fmt.Sprintf("view/%s/%s/%s.vue", a.Package, a.PackageName, a.PackageName)
// 	err = global.MAY_DB.Create(&menu).Error
// 	return menu.ID, err
// }

func (autoCodeService *AutoCodeService) getNeedList(autoCode *system.AutoCodeStruct) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	// 去空格
	utils.TrimSpace(autoCode)
	for _, filed := range autoCode.Fields {
		utils.TrimSpace(filed)
	}

	tplFiles, err := autoCodeService.GetAllTplFile(autocodePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFiles))
	fileList = make([]string, 0, len(tplFiles))
	needMkdir = make([]string, 0, len(tplFiles))
	// 根据文件路径生成 tplData 结构体
	for _, file := range tplFiles {
		dataList = append(dataList, tplData{locationPath: file, autoPackage: autoCode.Package})
	}

	for index, value := range dataList {
		dataList[index].template, err = template.ParseFiles(value.locationPath)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	// 生成文件路径，填充 autoCodePath 字段，readme.txt.tpl不符合规则，需要特殊处理
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt

	// resource/autocode_template/readme.txt.tpl
	// resource/autocode_template/server/**
	for index, value := range dataList {
		// 例如 value = resource/autocode_template/server/api.go.tpl
		// 例如 value = resource/autocode_template/web/api.js.tpl
		trimBase := strings.TrimPrefix(value.locationPath, autocodePath+"/")
		// if trimBase == "readme.txt.tpl" {
		// 	dataList[index].autoCodePath = autoPath + "readme.txt"
		// 	continue
		// }

		// trimBase = server/api.go.tpl
		// trimBase = web/api.js.tpl
		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			// origFileName = api.go
			// origFileName = api.js
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				var fileName string
				// origFileName[firstDot:] = go
				// origFileName[firstDot:] = js
				if origFileName[firstDot:] != ".go" {
					// fileName = admin.js
					fileName = autoCode.PackageName + origFileName[firstDot:]
				} else {
					// fileName = admin.go
					fileName = autoCode.HumpPackageName + origFileName[firstDot:]
				}
				// autocode_template/server/admin/api/admin.go
				// autocode_template/web/admin/api/admin.js
				dataList[index].autoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator], autoCode.PackageName,
					origFileName[:firstDot], fileName)
			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].autoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].autoCodePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

func injectionCode(structName string, bf *strings.Builder) error {
	for _, meta := range injectionPaths {
		code := fmt.Sprintf(meta.structNameF, structName)
		ast.ImportForAutoEnter(meta.path, meta.funcName, code)
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}

func (autoCodeService *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(pathName)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			fileList, err = autoCodeService.GetAllTplFile(pathName+"/"+file.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(file.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+file.Name())
			}
		}
	}
	return fileList, err
}
