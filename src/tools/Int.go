package tools

// SliceStringContainsElement 字符串序列是否包含
func SliceIntContainsElement(arr []int, ele int) bool {
	for _, v := range arr {
		if ele == v {
			return true
		}
	}
	return false
}
