package {{.Package}}

import (
	"github.com/shansec/go-vue-admin/api/v1"
	"github.com/shansec/go-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type {{.StructName}}Router struct {
}

// Init{{.StructName}}Router 初始化 {{.Description}} 路由信息
func (s *{{.StructName}}Router) Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.Abbreviation}}Record := Router.Group("{{.Abbreviation}}")
	{{.Abbreviation}}Api := v1.ApiGroupApp.{{.PackageT}}ApiGroup.{{.StructName}}Api
	{
		{{.Abbreviation}}Router.POST("create{{.StructName}}", {{.Abbreviation}}Api.Create{{.StructName}})   // 新建{{.Description}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}", {{.Abbreviation}}Api.Delete{{.StructName}}) // 删除{{.Description}}
		{{.Abbreviation}}Router.DELETE("delete{{.StructName}}ByIds", {{.Abbreviation}}Api.Delete{{.StructName}}ByIds) // 批量删除{{.Description}}
		{{.Abbreviation}}Router.PUT("update{{.StructName}}", {{.Abbreviation}}Api.Update{{.StructName}})    // 更新{{.Description}}
		{{.Abbreviation}}RouterWithoutRecord.GET("find{{.StructName}}", {{.Abbreviation}}Api.Find{{.StructName}})        // 根据ID获取{{.Description}}
		{{.Abbreviation}}RouterWithoutRecord.GET("get{{.StructName}}List", {{.Abbreviation}}Api.Get{{.StructName}}List)  // 获取{{.Description}}列表
		{{.Abbreviation}}RouterWithoutAuth.GET("get{{.StructName}}Public", {{.Abbreviation}}Api.Get{{.StructName}}Public)  // 获取{{.Description}}列表
	}
}
