package request

import (
	"github.com/shansec/go-vue-admin/dao/common/request"
	"github.com/shansec/go-vue-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
