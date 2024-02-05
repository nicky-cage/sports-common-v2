package tools

import "strings"

// SliceStringContainsElement 字符串序列是否包含
func SliceStringContainsElement(arr []string, ele string) bool {
	ele = strings.Replace(ele, " ", "", -1)
	ele = strings.ToLower(ele)
	for _, v := range arr {
		if ele == v {
			return true
		}
	}
	return false
}

// SliceRemoveDuplicateElement 去除字符串slice中重复的元素
func SliceRemoveDuplicateElement(adders []string) []string {
	result := make([]string, 0, len(adders))
	temp := map[string]struct{}{}
	for _, item := range adders {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// InSlice 是否在数组当中
func InSlice[T int | int8 | uint8 | int32 | uint32 | int64 | uint64 | bool | string](list []T, element T) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}
