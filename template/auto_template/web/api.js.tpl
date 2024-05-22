import service from '@/utils/request'

// @Tags {{.StructName}}
// @Summary 创建{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "创建{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "创建{{.Description}}失败"
// @Router /{{.Abbreviation}}/create{{.StructName}} [POST]
export const create{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/create{{.StructName}}',
    method: 'POST',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 删除{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "删除{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "删除{{.Description}}失败"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [DELETE]
export const delete{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}',
    method: 'DELETE',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 批量删除{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "批量删除{{.Description}}失败"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [DELETE]
export const delete{{.StructName}}ByIds = (params) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}ByIds',
    method: 'DELETE',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 更新{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "更新{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "更新{{.Description}}失败"
// @Router /{{.Abbreviation}}/update{{.StructName}} [PUT]
export const update{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/update{{.StructName}}',
    method: 'PUT',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 用id查询{{.Description}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.{{.StructName}} true "用id查询{{.Description}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "用id查询{{.Description}}失败"
// @Router /{{.Abbreviation}}/find{{.StructName}} [GET]
export const find{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/find{{.StructName}}',
    method: 'GET',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 分页获取{{.Description}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取{{.Description}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Failure 400 {object} response.Response "请求参数验证失败"
// @Failure 500 {object} response.Response "分页获取{{.Description}}列表失败"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [GET]
export const get{{.StructName}}List = (params) => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}List',
    method: 'GET',
    params
  })
}
