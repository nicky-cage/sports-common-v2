package validation

import (
	"regexp"
)

// 是否是电子邮件
func IsMail(mail string) bool {
	regString := `[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+`
	if matched, err := regexp.MatchString(regString, mail); !matched || err != nil {
		return false
	}
	return true
}

// 是否是用户名称
func IsUserName(userName string) bool {
	regString := `[A-Za-z]{1}[A-Za-z0-9_]{4,17}`
	if matched, err := regexp.MatchString(regString, userName); !matched || err != nil {
		return false
	}
	return true
}

// 是否是银行账号
func IsBankCard(cardNumber string) bool {
	regString := `[1-9]{1}[0-9]{14,18}`
	if matched, err := regexp.MatchString(regString, cardNumber); !matched || err != nil {
		return false
	}
	return true
}

// 是否是设备编号
func IsDeviceNumber(deviceNumber string) bool {
	regString := `[0-9a-zA-Z]{10,}`
	if matched, err := regexp.MatchString(regString, deviceNumber); !matched || err != nil {
		return false
	}
	return true
}

// 是否是手机号码
func IsPhoneNumber(phone string) bool {
	regString := `1[356789]{1}\d{9}`
	if matched, err := regexp.MatchString(regString, phone); !matched || err != nil {
		return false
	}
	return true
}
