package system

import (
	systemRes "github/shansec/go-vue-admin/dao/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"

	"github/shansec/go-vue-admin/dao/common/response"
	"github/shansec/go-vue-admin/global"
)

var store = base64Captcha.DefaultMemStore

// Captcha
// @Summary 获取验证码
// @Description 生成并返回一个新的默认数字验证码
// @Tags SysUser
// @Produce json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse, msg=string}	"获取验证码"
// @Failure 500 {object} response.Response "验证码获取失败"
// @Router /base/captcha [GET]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 生成默认数字验证码驱动
	driver := base64Captcha.NewDriverDigit(
		global.MAY_CONFIG.Captcha.ImgHeight,
		global.MAY_CONFIG.Captcha.ImgWidth,
		global.MAY_CONFIG.Captcha.KeyLong,
		0.7,
		80,
	)

	// 使用指定驱动和存储创建验证码实例
	cp := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	if id, b64s, _, err := cp.Generate(); err != nil {
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
