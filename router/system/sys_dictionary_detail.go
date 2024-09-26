package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shansec/go-vue-admin/api/v1"
)

type DictionaryDetailRouter struct{}

func (d *DictionaryDetailRouter) InitDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("/dictionaryDetail")
	dictionaryDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.POST("/createDictionaryDetail", dictionaryDetailApi.CreateDictionaryDetail)
		dictionaryDetailRouter.DELETE("/deleteDictionaryDetail", dictionaryDetailApi.DeleteDictionaryDetail)
		dictionaryDetailRouter.PUT("/updateDictionaryDetail", dictionaryDetailApi.UpdateDictionaryDetail)
		dictionaryDetailRouter.POST("/getSysDictionaryDetail", dictionaryDetailApi.GetSysDictionaryDetail)
		dictionaryDetailRouter.POST("/getDictionaryDetailList", dictionaryDetailApi.GetDictionaryDetailList)
	}
}
