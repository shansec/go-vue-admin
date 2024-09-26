package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shansec/go-vue-admin/api/v1"
)

type DictionaryRouter struct{}

func (d *DictionaryRouter) InitDictionaryRouter(Router *gin.RouterGroup) {
	dictionaryRouter := Router.Group("/dictionary")
	dictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		dictionaryRouter.POST("/createDictionary", dictionaryApi.CreateDictionary)
		dictionaryRouter.DELETE("/deleteDictionary", dictionaryApi.DeleteDictionary)
		dictionaryRouter.PUT("/updateDictionary", dictionaryApi.UpdateDictionary)
		dictionaryRouter.POST("/getDictionary", dictionaryApi.GetDictionary)
		dictionaryRouter.POST("/getDictionaryInfoList", dictionaryApi.GetDictionaryInfoList)
	}
}
