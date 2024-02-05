package request

import (
	"encoding/json"
	"sports-common/log"
	"sports-common/tools"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"xorm.io/builder"
)

// GetPostedData 获取提交的post来的数组
func GetPostedData(c *gin.Context) map[string]interface{} {
	if data, exists := c.Get("__posted_data"); exists {
		return data.(map[string]interface{})
	}

	if c == nil || c.Request == nil {
		log.Logger.Error("没有获取到任何内容: ")
		return nil
	}

	bytes, err := c.GetRawData()
	if err != nil {
		log.Err("获取POST原始数据有误: %v\n", err)
		return nil
	}
	if len(bytes) == 0 {
		return nil
	}

	var data = map[string]interface{}{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		kvs := strings.Split(string(bytes), "&")
		if len(kvs) == 0 {
			log.Err("获取POST并转换时失败: %v\n", err)
		}
		for _, v := range kvs {
			arr := strings.Split(v, "=")
			if len(arr) != 2 {
				continue
			}
			data[arr[0]] = arr[1]
		}
	}

	c.Set("__posted_data", data)
	return data
}

// GetPostAdminData 得到管理后台提交的post数据
func GetPostAdminData(c *gin.Context) string {
	bytes, err := c.GetRawData()
	if err != nil {
		log.Err("获取POST原始数据有误: %v\n", err)
		return ""
	}
	var data = map[string]interface{}{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Err("获取POST并转换时失败: %v\n", err)
	}
	b, err := json.Marshal(&data)
	if err != nil {
		log.Err("获取POST并转换json字符串时失败: %v\n", err)
		return ""
	}
	return string(b)
}

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Err("格式化获取页码出错: %v\n", err)
		return 1
	}
	if page == 0 {
		return 1
	}
	return page
}

// GetLimit 获取limit信息
func GetLimit(c *gin.Context) int {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "15"))
	if err != nil {
		log.Err("格式化获取 limit 出错: %v\n", err)
		return 15
	}
	if limit == 0 || limit > 10000 {
		return 15
	}
	return limit
}

// GetOffsets 获取 (limit, offset)
func GetOffsets(c *gin.Context) (int, int) {
	page := GetPage(c)
	size := GetLimit(c)
	return size, (page - 1) * size
}

// GetEsQueryCond 获取生成es相关条件
func GetEsQueryCond(c *gin.Context, data map[string]interface{}) *elastic.BoolQuery {
	cond := elastic.NewBoolQuery()
	//  boolQuery.Filter(elastic.NewTermQuery("game_type", gameType))
	//  boolQuery.Filter(elastic.NewRangeQuery("created_at").Gte(req.TimestampStart).Lt(req.TimestampEnd))
	for k, s := range data {
		val, exists := c.GetQuery(k)
		if !exists || val == "" {
			continue
		}
		switch sign := s.(type) {
		case string:
			switch sign {
			case "=":
				cond = cond.Must(elastic.NewTermQuery(k, val))
			case ">":
				cond = cond.Filter(elastic.NewRangeQuery(k).Gt(val))
			case ">=":
				cond = cond.Filter(elastic.NewRangeQuery(k).Gte(val))
			case "<":
				cond = cond.Filter(elastic.NewRangeQuery(k).Lt(val))
			case "<=":
				cond = cond.Filter(elastic.NewRangeQuery(k).Lte(val))
			case "%":
				cond = cond.Filter(elastic.NewQueryStringQuery(val).Field(k))
			case "between", "[]":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Gte(val))
				}, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Lte(val))
				})
			case "between|timestamp":
				setQueryCondBetween(c, k, func(val interface{}) {
					//tools.GetTimeStampByString(val.(string))
					cond = cond.Filter(elastic.NewRangeQuery(k).Gte(tools.GetTimeStampByString(val.(string))))
				}, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Lte(tools.GetTimeStampByString(val.(string))))
				})
			case "[]|timestamp":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Gte(tools.GetTimeStampByString(val.(string))))
				}, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Lte(tools.GetTimeStampByString(val.(string))))
				})
			case "()":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Gt(val))
				}, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Lt(val))
				})
			case "[)":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Gte(val))
				}, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Lt(val))
				})
			case "(]":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Gt(val))
				}, func(val interface{}) {
					cond = cond.Filter(elastic.NewRangeQuery(k).Lte(val))
				})
			default:
				continue
			}
		//case func(*gin.Context) builder.Cond: // 多余的条件
		//	cond = cond.And(s.(func(*gin.Context) builder.Cond)(c))
		default:
			continue
		}
	}
	return cond
}

// GetQueryCond 得到查询条件
func GetQueryCond(c *gin.Context, data map[string]interface{}) builder.Cond {
	cond := builder.NewCond()
	for k, s := range data {
		val, exists := c.GetQuery(k)
		if !exists || val == "" {
			continue
		}
		switch sign := s.(type) {
		case string:
			switch sign {
			case "=":
				cond = cond.And(builder.Eq{k: val})
			case ">":
				cond = cond.And(builder.Gt{k: val})
			case ">=":
				cond = cond.And(builder.Gte{k: val})
			case "<":
				cond = cond.And(builder.Lt{k: val})
			case "<=":
				cond = cond.And(builder.Lte{k: val})
			case "%":
				cond = cond.And(builder.Like{k, val})
			case "between", "[]": // 2个字段
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.And(builder.Gte{k: val})
				}, func(val interface{}) {
					cond = cond.And(builder.Lte{k: val})
				})
			case "between|timestamp": // 2个字段
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.And(builder.Gte{k: tools.GetTimeStampByString(val.(string))})
				}, func(val interface{}) {
					cond = cond.And(builder.Lte{k: tools.GetTimeStampByString(val.(string))})
				})
			case "[]|timestamp": // 2个字段
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.And(builder.Gte{k: tools.GetTimeStampByString(val.(string))})
				}, func(val interface{}) {
					cond = cond.And(builder.Lte{k: tools.GetTimeStampByString(val.(string))})
				})
			case "[_]|timestamp": // 1个字段拆分 => 比如 created, 传过来的是时间区间 2020-01-02 - 2020-02-02 => 转换为时间戳对比
				var startAt int64
				var endAt int64
				value, exists := c.GetQuery(k)
				if exists {
					areas := strings.Split(value, " - ")
					startAt = tools.GetTimeStampByString(areas[0])
					endAt = tools.GetTimeStampByString(areas[1])
					cond = cond.And(builder.Gte{k: startAt}).And(builder.Lte{k: endAt})
				}
			case "()":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.And(builder.Gt{k: val})
				}, func(val interface{}) {
					cond = cond.And(builder.Lt{k: val})
				})
			case "[)":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.And(builder.Gte{k: val})
				}, func(val interface{}) {
					cond = cond.And(builder.Lt{k: val})
				})
			case "(]":
				setQueryCondBetween(c, k, func(val interface{}) {
					cond = cond.And(builder.Gt{k: val})
				}, func(val interface{}) {
					cond = cond.And(builder.Lte{k: val})
				})
			default:
				continue
			}
		case func(*gin.Context) builder.Cond: // 多余的条件
			cond = cond.And(s.(func(*gin.Context) builder.Cond)(c))
		default:
			continue
		}
	}
	return cond
}

// setQueryCondBetween 设置以区间为条件的查询字段
func setQueryCondBetween(c *gin.Context, field string, callStart func(interface{}), callEnd func(interface{})) {
	start := field + "_start"
	end := field + "_end"
	if startVal, exists := c.GetQuery(start); exists {
		callStart(startVal)
	}
	if endVal, exists := c.GetQuery(end); exists {
		callEnd(endVal)
	}
}

// QueryCondEq 批量处理查询条件 - eq
func QueryCondEq(c *gin.Context, cond *builder.Cond, fields map[string]string) {
	for queryField, field := range fields {
		if val := c.DefaultQuery(queryField, ""); val != "" {
			*cond = (*cond).And(builder.Eq{field: val})
		}
	}
}

// QueryCondNeq 批量处理查询条件 - neq - 不等于
func QueryCondNeq(c *gin.Context, cond *builder.Cond, fields map[string]string) {
	for queryField, field := range fields {
		if val := c.DefaultQuery(queryField, ""); val != "" {
			*cond = (*cond).And(builder.Neq{field: val})
		}
	}
}

// QueryCondGte 批量处理查询条件 - gte 大于等于
func QueryCondGte(c *gin.Context, cond *builder.Cond, fields map[string]string) {
	for queryField, field := range fields {
		if val := c.DefaultQuery(queryField, ""); val != "" {
			*cond = (*cond).And(builder.Gte{field: val})
		}
	}
}

// QueryCondLte 批量处理查询条件 - lte 小于等于
func QueryCondLte(c *gin.Context, cond *builder.Cond, fields map[string]string) {
	for queryField, field := range fields {
		if val := c.DefaultQuery(queryField, ""); val != "" {
			*cond = (*cond).And(builder.Lte{field: val})
		}
	}
}

// QueryCondLike 批量处理查询条件 - like
func QueryCondLike(c *gin.Context, cond *builder.Cond, fields map[string]string) {
	for queryField, field := range fields {
		if val := c.DefaultQuery(queryField, ""); val != "" {
			*cond = (*cond).And(builder.Like{field, val})
		}
	}
}

// QueryCond 批量处理 like/=
// conds => map[操作符]map[查询字段名称]数据库字段名称
func QueryCond(c *gin.Context, cond *builder.Cond, conds map[string]map[string]string) {
	for opt, fields := range conds {
		if opt == "%" {
			QueryCondLike(c, cond, fields)
		} else if opt == "=" {
			QueryCondEq(c, cond, fields)
		} else if opt == ">=" {
			QueryCondGte(c, cond, fields)
		} else if opt == "<=" {
			QueryCondLte(c, cond, fields)
		} else if opt == "!=" {
			QueryCondNeq(c, cond, fields)
		}
	}
}
