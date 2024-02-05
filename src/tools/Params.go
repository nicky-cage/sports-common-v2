package tools

import (
	"regexp"
	"sports-common/consts"
	"strconv"
	"strings"
	"time"
)

// CheckDateFormat 校验日期格式
func CheckDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// CheckRealName 校验真实姓名
func CheckRealName(realName string) bool {
	ok, _ := regexp.MatchString("^[\u4E00-\u9FA5]{2,10}$", realName)
	return ok
}

// CheckBankBranch 校验银行支行
func CheckBankBranch(branch string) bool {
	ok, _ := regexp.MatchString("^[\u4E00-\u9FA5]{4,20}$", branch)
	return ok
}

// CheckBankAddress 校验银行支行
func CheckBankAddress(address string) bool {
	ok, _ := regexp.MatchString("^[\u4e00-\u9fa5_a-zA-Z0-9]{4,20}$", address)
	return ok
}

// CheckUserName 校验用户名  4-11位字母或数字
func CheckUserName(userName string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,12}$", userName)
	return ok
}

// CheckPassword 校验密码 6-12位字母或数字
//密码md5 32位
func CheckPassword(password string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{32}$", password)
	return ok
}

// CheckVCode 校验手机邮箱验证码 只允许输入4-6位字母或数字
func CheckVCode(code string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,6}$", code)
	return ok
}

// CheckCVCode 图形验证码
func CheckCVCode(code string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,6}$", code)
	return ok
}

// CheckCVID 图形验证码ID
func CheckCVID(id string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9]{1,100}$", id)
	return ok
}

// FormatPhoneNumber 格式化电话号码
func FormatPhoneNumber(number string) string {
	runeNumber := []rune(number)
	if len(runeNumber) != 11 {
		return ""
	}
	front := string(runeNumber[:3])
	tail := string(runeNumber[7:])
	return front + "****" + tail
}

// FormatEmail 格式化邮件
func FormatEmail(email string) string {
	ss := strings.Split(email, "@")
	if len(ss) != 2 {
		return ""
	}

	var front string
	if len(ss[0]) < 2 {
		front = ss[0]
	} else {
		front = string([]rune(ss[0])[:2])
	}

	return front + "****@" + ss[1]
}

// CheckPhoneNumber 校验手机号
func CheckPhoneNumber(phone string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0-8])|(18[0-9])|166|198|(19[1-9])|(147))\\d{8}$"
	ok, _ := regexp.MatchString(pattern, phone)
	return ok
}

// CheckEmail 校验邮箱
func CheckEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	ok, _ := regexp.MatchString(pattern, email)
	return ok
}

// CheckBankCard 校验卡号
func CheckBankCard(bankCard string) bool {
	pattern := `^[0-9]{16,20}$`
	ok, _ := regexp.MatchString(pattern, bankCard)
	return ok
}

// CheckBankCode 检查银行编码
func CheckBankCode(bankCode string) bool {
	for _, v := range consts.BankList {
		if v.BankCode == bankCode {
			return true
		}
	}
	return false
}

// CheckDeviceID 设备ID
func CheckDeviceID(deviceID string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9:\\-]{1,100}$", deviceID)
	return ok
}

// CheckHttpUrl 检查url
func CheckHttpUrl(url string) bool {
	ok, _ := regexp.MatchString(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`, url)
	return ok
}

// CheckIP 检查ip
func CheckIP(ip string) bool {
	reg := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	ok, _ := regexp.MatchString(reg, ip)
	return ok
}

// CheckVType 检查vtype
func CheckVType(vType string) bool {
	vt, _ := strconv.Atoi(vType)
	if vt < 1 || vt > 20 {
		return false
	}
	return true
}

// CheckDirFormat 检查dir格式
func CheckDirFormat(dir string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9/.]{1,100}$", dir)
	return ok
}
