package utils

import (
	"fmt"
	"strings"
)

// IfString 依据条件返回
func IfString(c bool, s1, s2 string) string {
	if c {
		return s1
	}
	return s2
}

// MapToInsertSQL 转换成SQL语句
func MapToInsertSQL(tableName string, data map[string]interface{}) []interface{} {
	fArr := []string{}
	vArr := []string{}
	rArr := []interface{}{}

	for k, v := range data {
		fArr = append(fArr, k)
		vArr = append(vArr, "?")
		rArr = append(rArr, v)
	}

	rSQL := "INSERT INTO " + tableName + " (" + strings.Join(fArr, ", ") + ") VALUES (" + strings.Join(vArr, ", ") + ")"
	rList := []interface{}{rSQL}
	rList = append(rList, rArr...)
	return rList
}

// MapToUpdateSQL 转换成SQL语句
func MapToUpdateSQL(tableName string, data map[string]interface{}, whCond string) []interface{} {
	fArr := []string{}
	rArr := []interface{}{}

	for k, v := range data {
		fArr = append(fArr, fmt.Sprintf("%s = ?", k))
		rArr = append(rArr, v)
	}

	rSQL := "UPDATE " + tableName + " SET " + strings.Join(fArr, ", ") + " WHERE " + whCond
	rList := []interface{}{rSQL}
	rList = append(rList, rArr...)
	return rList
}
