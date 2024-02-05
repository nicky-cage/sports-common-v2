package pgsql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// PgQueryBuilder 查询构建
type PgQueryBuilder struct {
	Query      []string
	Conditions []interface{}
}

// NewQueryBuilder 构建查询
func NewQueryBuilder() *PgQueryBuilder {
	return &PgQueryBuilder{}
}

// Eq 等于
func (ths *PgQueryBuilder) Eq(k string, v interface{}) *PgQueryBuilder {
	ths.Query = append(ths.Query, k+" = ?")
	ths.Conditions = append(ths.Conditions, v)
	return ths
}

// Lte 小于等于
func (ths *PgQueryBuilder) Lte(k string, v interface{}) *PgQueryBuilder {
	ths.Query = append(ths.Query, k+" <= ?")
	ths.Conditions = append(ths.Conditions, v)
	return ths
}

// Lt 小于
func (ths *PgQueryBuilder) Lt(k string, v interface{}) *PgQueryBuilder {
	ths.Query = append(ths.Query, k+" < ?")
	ths.Conditions = append(ths.Conditions, v)
	return ths
}

// Gte 大于等于
func (ths *PgQueryBuilder) Gte(k string, v interface{}) *PgQueryBuilder {
	ths.Query = append(ths.Query, k+" >= ?")
	ths.Conditions = append(ths.Conditions, v)
	return ths
}

// Gt 大于
func (ths *PgQueryBuilder) Gt(k string, v interface{}) *PgQueryBuilder {
	ths.Query = append(ths.Query, k+" > ?")
	ths.Conditions = append(ths.Conditions, v)
	return ths
}

// Like 模糊
func (ths *PgQueryBuilder) Like(k string, v interface{}) *PgQueryBuilder {
	ths.Query = append(ths.Query, k+" LIKE '%%?%%'")
	ths.Conditions = append(ths.Conditions, v)
	return ths
}

// In 范围
func (ths *PgQueryBuilder) In(conds map[string]string) *PgQueryBuilder {
	for k, v := range conds {
		ths.Query = append(ths.Query, k+" LIKE %?%")
		ths.Conditions = append(ths.Conditions, v)
	}
	return ths
}

// Build 构造查询
func (ths *PgQueryBuilder) Build() (string, []interface{}) {
	return strings.Join(ths.Query, " AND "), ths.Conditions
}

// Queries 依据查询条件
func (ths *PgQueryBuilder) Queries(c *gin.Context, data map[string]string) *PgQueryBuilder {
	for key, field := range data {
		if val, exists := c.GetQuery(key); exists {
			ths.Query = append(ths.Query, field+" = ?")
			ths.Conditions = append(ths.Conditions, val)
		}
	}
	return ths
}

// QueriesByInt 依据查询条件
func (ths *PgQueryBuilder) QueriesByInt(c *gin.Context, data map[string]string) *PgQueryBuilder {
	for key, field := range data {
		if val, exists := c.GetQuery(key); exists {
			ths.Query = append(ths.Query, field+" = ?")
			v, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("转换条件失败:", field, key, val)
				continue
			}
			ths.Conditions = append(ths.Conditions, v)
		}
	}
	return ths
}
