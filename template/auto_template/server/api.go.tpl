package {{.Package}}

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"

		"github/shansec/go-vue-admin/global"
    "github/shansec/go-vue-admin/model/{{.Package}}"
    {{.Package}}Req "github/shansec/go-vue-admin/model/{{.Package}}/request"
    "github/shansec/go-vue-admin/model/common/response"
    "github/shansec/go-vue-admin/service"
    {{- if .AutoCreateResource}}
    "github/shansec/go-vue-admin/utils"
    {{- end }}
)

type {{.StructName}}Api struct {
}

var {{.Abbreviation}}Service = service.ServiceGroupApp.{{.PackageT}}ServiceGroup.{{.StructName}}Service


// Create{{.StructName}} 创建{{.Description}}
// @Tags {{.StructName}}
// @Summary 创建{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "创建{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "创建{{.Description}}失败"
// @Router /{{.Abbreviation}}/create{{.StructName}} [POST]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
    {{.Abbreviation}}.CreatedBy = utils.GetUserID(c)
	{{- end }}

	if err := {{.Abbreviation}}Service.Create{{.StructName}}(&{{.Abbreviation}}); err != nil {
    global.MAY_LOGGER.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete{{.StructName}} 删除{{.Description}}
// @Tags {{.StructName}}
// @Summary 删除{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "删除{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "删除{{.Description}}失败"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [DELETE]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}(c *gin.Context) {
	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
		{{- if .AutoCreateResource }}
    	userID := utils.GetUserID(c)
    {{- end }}
	if err := {{.Abbreviation}}Service.Delete{{.StructName}}({{.PrimaryField.FieldJson}} {{- if .AutoCreateResource -}},userID{{- end -}}); err != nil {
    global.MAY_LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Delete{{.StructName}}ByIds 批量删除{{.Description}}
// @Tags {{.StructName}}
// @Summary 批量删除{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "批量删除{{.Description}}失败"
// @Router /{{.Abbreviation}}/delete{{.StructName}}ByIds [DELETE]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}ByIds(c *gin.Context) {
	{{.PrimaryField.FieldJson}}s := c.QueryArray("{{.PrimaryField.FieldJson}}s[]")
    	{{- if .AutoCreateResource }}
    userID := utils.GetUserID(c)
        {{- end }}
	if err := {{.Abbreviation}}Service.Delete{{.StructName}}ByIds({{.PrimaryField.FieldJson}}s{{- if .AutoCreateResource }},userID{{- end }}); err != nil {
    global.MAY_LOGGER.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update{{.StructName}} 更新{{.Description}}
// @Tags {{.StructName}}
// @Summary 更新{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.Package}}.{{.StructName}} true "更新{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "更新{{.Description}}失败"
// @Router /{{.Abbreviation}}/update{{.StructName}} [PUT]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	    {{- if .AutoCreateResource }}
    {{.Abbreviation}}.UpdatedBy = utils.GetUserID(c)
        {{- end }}

	if err := {{.Abbreviation}}Service.Update{{.StructName}}({{.Abbreviation}}); err != nil {
    global.MAY_LOGGER.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find{{.StructName}} 用id查询{{.Description}}
// @Tags {{.StructName}}
// @Summary 用id查询{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}.{{.StructName}} true "用id查询{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "用id查询{{.Description}}失败"
// @Router /{{.Abbreviation}}/find{{.StructName}} [GET]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Find{{.StructName}}(c *gin.Context) {
	{{.PrimaryField.FieldJson}} := c.Query("{{.PrimaryField.FieldJson}}")
	if re{{.Abbreviation}}, err := {{.Abbreviation}}Service.Get{{.StructName}}({{.PrimaryField.FieldJson}}); err != nil {
    global.MAY_LOGGER.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"re{{.Abbreviation}}": re{{.Abbreviation}}}, c)
	}
}

// Get{{.StructName}}List 分页获取{{.Description}}列表
// @Tags {{.StructName}}
// @Summary 分页获取{{.Description}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "分页获取{{.Description}}列表失败"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [GET]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
	var pageInfo {{.Package}}Req.{{.StructName}}Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
      response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// Get{{.StructName}}Public 不需要鉴权的{{.Description}}接口
// @Tags {{.StructName}}
// @Summary 不需要鉴权的{{.Description}}接口
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [GET]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}Public(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的{{.Description}}接口信息",
    }, "获取成功", c)
}
