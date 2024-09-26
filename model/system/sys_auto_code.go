package system

import (
	"go/token"
	"strings"

	"github.com/shansec/go-vue-admin/global"
)

type SysAutoCode struct {
	global.MAY_MODEL
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:标签"`
	Desc        string `json:"desc" gorm:"comment:描述"`
}

func (SysAutoCode) TableName() string {
	return "sys_auto_codes"
}

type AutoCodeStruct struct {
	StructName          string   `json:"structName"`          // Struct名称
	TableName           string   `json:"tableName"`           // 表名
	PackageName         string   `json:"packageName"`         // 文件名称
	HumpPackageName     string   `json:"humpPackageName"`     // go文件名称
	Abbreviation        string   `json:"abbreviation"`        // Struct简称
	Description         string   `json:"description"`         // Struct中文名称
	AutoCreateApiToSql  bool     `json:"autoCreateApiToSql"`  // 是否自动创建api
	AutoCreateMenuToSql bool     `json:"autoCreateMenuToSql"` // 是否自动创建menu
	AutoCreateResource  bool     `json:"autoCreateResource"`  // 是否自动创建资源标识
	AutoMoveFile        bool     `json:"autoMoveFile"`        // 是否自动移动文件
	BusinessDB          string   `json:"businessDB"`          // 业务数据库
	DefaultModel        bool     `json:"defaultModel"`        // 是否使用默认Model
	Fields              []*Field `json:"fields"`
	PrimaryField        *Field   `json:"primaryField"`
	HasTimer            bool     `json:"-"`
	HasSearchTimer      bool     `json:"-"`
	DictTypes           []string `json:"-"`
	Package             string   `json:"package"`
	PackageT            string   `json:"-"`
	NeedSort            bool     `json:"-"`
	HasPic              bool     `json:"-"`
	HasRichText         bool     `json:"-"`
	HasFile             bool     `json:"-"`
	NeedJSON            bool     `json:"-"`
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	DictType        string `json:"dictType"`        // 字典
	Require         bool   `json:"require"`         // 是否必填
	ErrorText       string `json:"errorText"`       // 校验失败文字
	Clearable       bool   `json:"clearable"`       // 是否可清空
	Sort            bool   `json:"sort"`            // 是否增加排序
	PrimaryKey      bool   `json:"primaryKey"`      // 是否主键
}

func (a *AutoCodeStruct) Pretreatment() {
	a.KeyWord()
	a.SuffixTest()
}

// KeyWord 是go关键字的处理加上 _ ，防止编译报错
func (a *AutoCodeStruct) KeyWord() {
	if token.IsKeyword(a.Abbreviation) {
		a.Abbreviation = a.Abbreviation + "_"
	}
}

// SuffixTest 处理_test 后缀
func (a *AutoCodeStruct) SuffixTest() {
	if strings.HasSuffix(a.HumpPackageName, "test") {
		a.HumpPackageName = a.HumpPackageName + "_"
	}
}
