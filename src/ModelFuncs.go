package common

import (
	"sports-common/config"
	"sports-common/log"
)

// 获取提交上来的post数组
func getPostedFieldValues(platform string, tableName string, isCreate bool, postedData map[string]interface{}) (map[string]interface{}, []string, []interface{}) {
	realFields := config.GetTableFields(platform, tableName) //获取对应的表信息
	currentTime := Now()                                     //当前时间
	var (
		fields []string      //字段
		values []interface{} //写入的值
	)
	for _, fv := range realFields {
		if fv == "id" {
			continue
		}
		if fv == "updated" || (fv == "created" && isCreate) {
			fields = append(fields, fv)
			values = append(values, currentTime)
			continue
		}
		if val, exists := postedData[fv]; exists { //&& fmt.Sprintf("%v", val) != "" { //如果确实有提交字段信息
			fields = append(fields, fv)
			values = append(values, val)
		}
	}
	if len(fields) > 0 {
		return postedData, fields, values
	}
	log.Err("没有提交任何可以与数据库匹配的字段: %v\n%v\n", postedData, realFields)
	return nil, nil, nil
}
