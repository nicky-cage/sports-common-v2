package utils

import "fmt"

// ToStringSlice 转换为字符串序列
func ToStringSlice[T int8 | uint8 | int16 | uint16 | int32 | uint32 | int | uint | int64 | uint64](arr []T) []string {
	rArr := []string{}
	for _, r := range arr {
		rArr = append(rArr, fmt.Sprintf("%d", r))
	}
	return rArr
}
