package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/model/common/response"
	systemRes "github/shansec/go-vue-admin/model/system/response"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func (b *BaseApi) Captcha(c *gin.Context) {
	// 生成默认数字
	driver := base64Captcha.NewDriverDigit(global.MAY_CONFIG.Captcha.ImgHeight, global.MAY_CONFIG.Captcha.ImgWidth, global.MAY_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.MAY_LOGGER.Error("验证码获取失败", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.MAY_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", c)
	}
}
