package config

import (
	"strings"

	"gopkg.in/ini.v1"
)

// 依据key获得ini的设置项
func getIniKey(key string) *ini.Key {
	sArr := strings.Split(key, ".")
	arrLen := len(sArr)
	if arrLen == 1 {
		return Ini.Section("").Key(key)
	} else if arrLen == 2 {
		return Ini.Section(sArr[0]).Key(sArr[1])
	}
	return nil
}
