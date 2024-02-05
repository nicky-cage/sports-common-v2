package captchas

import (
	"strings"

	"github.com/mojocn/base64Captcha"
)

// CaptchaSource 默认生成图片Source
var CaptchaSource = "0123456789"

// Captcha 生成验证码
type Captcha struct {
	Driver base64Captcha.Driver
	Store  base64Captcha.Store
}

// Capt 生成默认验主码
var Capt = func(platform string) *Captcha {
	return DefaultCaptcha(platform)
}

// GenerateCaptcha 生成验证码
func (p *Captcha) GenerateCaptcha() map[string]interface{} {
	c := base64Captcha.NewCaptcha(p.Driver, p.Store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	return body
}

// Verify 校验验证码
func (p *Captcha) Verify(id, verifyValue string) bool {
	//true  删除验证码
	verifyValue = strings.ToLower(verifyValue)
	return p.Store.Verify(id, verifyValue, true)
}
