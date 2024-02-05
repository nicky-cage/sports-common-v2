package captchas

import "github.com/mojocn/base64Captcha"

// NewCaptcha 生成指定宽高的验证码
func NewCaptcha(platform string, height, width, length int, source string) *Captcha {
	driverStr := &base64Captcha.DriverString{
		Height: height,
		Width:  width,
		Source: source,
		Length: length,
	}
	cp := &Captcha{
		Store:  &StoreRedis{Platform: platform},
		Driver: driverStr.ConvertFonts(),
	}
	return cp
}

// NewCaptchaWith 生成指定宽高验证码
func NewCaptchaWith(platform string, height, width, length int) *Captcha {
	driverStr := &base64Captcha.DriverString{
		Height: height,
		Width:  width,
		Source: CaptchaSource,
		Length: length,
	}
	cp := &Captcha{
		Store:  &StoreRedis{Platform: platform},
		Driver: driverStr.ConvertFonts(),
	}
	return cp
}

// DefaultCaptcha 生成默认验证码
func DefaultCaptcha(platform string) *Captcha {
	driverStr := &base64Captcha.DriverString{
		Height: 50,
		Width:  100,
		Source: "0123456789", //qwertyuiopasdfghjklzxcvbnm
		Length: 4,
	}
	cp := &Captcha{
		Store:  &StoreRedis{Platform: platform},
		Driver: driverStr.ConvertFonts(),
	}
	return cp
}
