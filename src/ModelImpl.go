package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"sports-common/consts"
	"sports-common/db"
	"sports-common/log"
	"sports-common/tools"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"xorm.io/builder"
)

// GetTableName 获取表名称
func (p *Model) GetTableName() string {
	return p.TabName
}

// GetFindParameters 获取分页等相关信息
func (p *Model) GetFindParameters(params ...interface{}) (builder.Cond, int, int, string) {
	var args []interface{}
	if realArgs, ok := params[0].([]interface{}); ok { //兼容多种格式
		args = realArgs
	} else {
		args = params
	}
	argLength := len(args)

	limit := 15
	offset := 0
	cond := builder.NewCond().And(builder.Eq{"1": 1})
	orderBy := ""

	if argLength >= 1 { //条件
		if args[0] != nil {
			cond = args[0].(builder.Cond)
		}
		if argLength >= 2 {
			limit = args[1].(int)
			if argLength >= 3 {
				offset = args[2].(int)
				if argLength >= 4 {
					orderBy = args[3].(string)
				}
			}
		}
	}
	return cond, limit, offset, orderBy
}

// FindAllNoCount 获取所有记录
// 参数
//	rows 对象数组
//	args: [查询条件, orderBy]
// 结果: 记录条数, 错误
// FindAllNoCount 依据平台识别号获取所有记录
func (p *Model) FindAllNoCount(platform string, rows interface{}, args ...interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	dbSession := p.GetSession(platform).Table(p.TabName)
	defer dbSession.Close()
	if len(p.Joins) >= 1 {
		for _, j := range p.Joins {
			dbSession.Join(j[0], j[1], j[2])
		}
	}
	argLength := len(args)
	if argLength >= 1 { //条件 - where
		if args[0] != nil {
			dbSession = dbSession.Where(args[0].(builder.Cond))
		}
	}
	if argLength >= 2 { // 排序 - order - by
		dbSession = dbSession.OrderBy(args[1].(string))
	}

	return dbSession.Find(rows)
}

// FindAll 获取所有记录
// 参数
//	rows 对象数组
//	args: [查询条件, limit, offset, orderBy]
// 结果: 记录条数, 错误
// FindAll 依据平台识别号获取所有记录
func (p *Model) FindAll(platform string, rows interface{}, args ...interface{}) (uint64, error) {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	session := p.GetSession(platform).Table(p.TabName)
	defer session.Close()
	if len(p.Joins) >= 1 {
		for _, j := range p.Joins {
			session.Join(j[0], j[1], j[2])
		}
	}
	orderBy := "id DESC"
	offset := 0
	cond := builder.NewCond()
	limit := 15
	if len(args) > 0 {
		cond, limit, offset, orderBy = p.GetFindParameters(args...)
	}
	if cond != nil && cond.IsValid() {
		session = session.Where(cond)
	}
	if orderBy != "" {
		session = session.OrderBy(orderBy)
	}
	total, err := session.Limit(limit, offset).FindAndCount(rows)
	return uint64(total), err
}

// Find 参数:
//	row: 对象
//	args: [查询条件]
// 返回: 是否存在记录, 错误
// Find 获取单条记录
func (p *Model) Find(platform string, row interface{}, args ...interface{}) (bool, error) {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	dbSession := p.GetSession(platform).Table(p.TabName)
	defer dbSession.Close()
	argLength := len(args)
	if argLength >= 1 { // 如果有第一个参数, 则是 WHERE => Cond.NewCond()
		if args[0] != nil {
			dbSession = dbSession.Where(args[0].(builder.Cond))
		}
	}
	if argLength >= 2 { //如果有第二个参数, 则是 ORDER_BY
		if args[1] != nil { // order-by
			dbSession = dbSession.OrderBy(args[1].(string))
		}
	}
	ok, err := dbSession.Get(row)
	return ok, err
}

// FindById 获取编号、平台获取记录
func (p *Model) FindById(platform string, id int, row interface{}) (bool, error) {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	dbSession := p.GetSession(platform).Table(p.TabName)
	defer dbSession.Close()
	cond := builder.And(builder.Eq{"id": id})
	dbSession.Where(cond)
	return dbSession.Get(row)
}

// CreateOne 添加一条
func (p *Model) CreateOne(platform string, m interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	_, err := p.GetSession(platform).InsertOne(m)
	return err
}

// CreateAll 依据平台识别号，一次性创建多条记录
func (p *Model) CreateAll(platform string, postedData []map[string]interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	var vals []string
	flds := ""
	inserts := []interface{}{""}
	// 先检测所有字段是否合法
	for _, v := range postedData {
		_, fields, values := getPostedFieldValues(platform, p.TabName, true, v)
		if fields == nil {
			return errors.New("批量添加失败: 无法依据平台获取字段信息")
		}
		if flds == "" {
			flds = fmt.Sprintf("(%s)", strings.Join(fields, ","))
		}
		vals = append(vals, fmt.Sprintf("(?%s)", strings.Repeat(",?", len(fields)-1)))
		inserts = append(inserts, values...)
	}
	inserts[0] = "INSERT INTO " + p.TabName + " " + flds + " VALUES " + strings.Join(vals, ",")
	dbSession := p.GetSession(platform)
	defer dbSession.Close()
	_, err := dbSession.Exec(inserts...)
	return err
}

// Create 依据平台识别号添加记录
func (p *Model) Create(platform string, postedData map[string]interface{}) (uint64, error) {
	_, fields, values := getPostedFieldValues(platform, p.TabName, true, postedData)
	if fields == nil {
		return 0, errors.New("添加数据失败: 无法依据平台获取字段信息")
	}
	inserts := []interface{}{""}
	inserts = append(inserts, values...)
	inserts[0] = fmt.Sprintf("INSERT INTO %s (%s) VALUES (?%s)", p.TabName, strings.Join(fields, ","), strings.Repeat(",?", len(fields)-1))
	dbSession := p.GetSession(platform)
	defer dbSession.Close()
	result, err := dbSession.Exec(inserts...)
	if err != nil {
		sql, _ := dbSession.LastSQL()
		log.Err("执行Model-Create时失败: %v\nSQL: %s\n%v\n", err, sql, inserts)
		return 0, errors.New("添加失败: " + err.Error())
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Err("获取添加记录时失败: %v\n", err)
		return 0, errors.New("添加失败: " + err.Error())
	}
	return uint64(lastInsertID), nil
}

// Update 修改记录
func (p *Model) Update(platform string, postedData map[string]interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	data, fields, values := getPostedFieldValues(platform, p.TabName, false, postedData)
	if fields == nil {
		return errors.New("修改失败: 无法获取字段信息")
	}
	idUn, exists := data["id"]
	if !exists {
		log.Err("修改记录时找不到ID\n")
		return errors.New("修改失败: 无法获取记录编号")
	}
	updates := []interface{}{""}
	updateFields := ""
	for k, v := range fields {
		updates = append(updates, values[k])
		updateFields += fmt.Sprintf("`%s` = ? ", v) + ","
	}
	id, _ := strconv.Atoi(fmt.Sprintf("%v", idUn))
	updates[0] = fmt.Sprintf("UPDATE %s SET %s WHERE id = %v", p.TabName, strings.Trim(updateFields, ","), strconv.Itoa(id))
	dbSession := p.GetSession(platform)
	defer dbSession.Close()
	if _, err := dbSession.Exec(updates...); err != nil {
		log.Err("执行Model-Update时失败: %v\nSQL: %s\n", err, updates)
		return errors.New("修改失败: " + err.Error())
	}
	return nil
}

// Delete 删除记录
func (p *Model) Delete(platform string, idStr string) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	var ids []string
	idStrArr := strings.Split(idStr, ",")
	for _, v := range idStrArr {
		_, err := strconv.Atoi(v)
		if err != nil {
			log.Err("错误编号: %s/%v\n", idStr, err)
			return errors.New("删除失败: " + err.Error())
		}
		ids = append(ids, v)
	}
	if len(ids) == 0 {
		log.Err("无效的ID: %s\n", idStr)
		return errors.New("删除失败: 记录编号无效")
	}
	sql := "DELETE FROM " + p.TabName + " WHERE id IN (" + strings.Join(ids, ",") + ")"
	dbSession := p.GetSession(platform)
	defer dbSession.Close()
	_, err := dbSession.Exec(sql)
	if err != nil {
		log.Err("删除失败: %v\n", err)
		return errors.New("删除失败: " + err.Error())
	}
	return nil
}

// FindByIdHash 依据idhash
func (p *Model) FindByIdHash(platform string, cacheKey string, id int, rClient *redis.Conn, tableModel interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	if p.CachePrefix != "" {
		cacheKey = p.CachePrefix + strconv.Itoa(id)
	}

	isHasInt, err := rClient.Exists(cacheKey).Result()
	if err == redis.Nil || isHasInt == 0 {
		cond := builder.NewCond().And(builder.Eq{"id": id})
		log.Logger.Info(tableModel)
		isExist, err := p.Find(platform, tableModel, cond)
		if err != nil {
			return err
		}
		if !isExist {
			msg := "记录不存在"
			log.Logger.Error(strconv.Itoa(id) + msg)
			return errors.New(msg)
		}
		//1. struct to json
		//2. json to map
		tbJson, err := json.Marshal(tableModel)
		if err != nil {
			log.Logger.Error(err.Error())
			return err
		}
		var m map[string]interface{}
		err = json.Unmarshal(tbJson, &m)
		if err != nil {
			log.Logger.Error(err.Error())
			return err
		}

		isOk, _ := rClient.HMSet(cacheKey, m).Result()
		if isOk { //7*24*60*60
			rClient.Expire(cacheKey, time.Minute*604800)
		}
	} else if err != nil {
		return err
	} else if isHasInt > 0 {
		tableModelStrMap, err := rClient.HGetAll(cacheKey).Result()
		if err != nil {
			return err
		}
		tools.SetStructFieldByJsonName(tableModel, tableModelStrMap)
	}
	return nil
}

// FindByIdKv isWrite 强制读数据库写缓存 比如操作account的钱后，立即调用此函数isWrite=true
func (p *Model) FindByIdKv(platform string, cacheKey string, id int, isWrite bool, rClient *redis.Conn, tableModel interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	// FindByIdKey 依据IdKey缓存
	if p.CachePrefix != "" {
		cacheKey = p.CachePrefix + strconv.Itoa(id)
	}
	intHas, err := rClient.Exists(cacheKey).Result()
	if err != nil && err != redis.Nil {
		return err
	} else if err == redis.Nil || intHas == 0 || isWrite {
		cond := builder.NewCond().And(builder.Eq{"id": id})
		isExist, err := p.Find(platform, tableModel, cond)
		if err != nil {
			return err
		}
		if !isExist {
			msg := "记录不存在"
			log.Logger.Error(strconv.Itoa(id) + msg)
			return errors.New(msg)
		}
		tableModelStrBytes, err := json.Marshal(tableModel)
		if err != nil {
			return err
		}
		rClient.Set(cacheKey, tableModelStrBytes, 120*time.Minute)
	} else {
		tableModelStr, err := rClient.Get(cacheKey).Result()
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(tableModelStr), tableModel)
		if err != nil {
			return err
		}
	}
	return nil
}

// FindByAndWhereKv 依据条件从缓存获取
func (p *Model) FindByAndWhereKv(platform string, cacheKey string, rClient *redis.Conn, tableModel interface{}, errorMsg string, whereArgs ...interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	if p.CachePrefix != "" {
		cacheKey = p.CachePrefix
	}
	isExists, err := rClient.Exists(cacheKey).Result()
	if err == redis.Nil || isExists == 0 {
		isExist, err := p.Find(platform, tableModel, whereArgs...)
		if err != nil {
			log.Logger.Error(errorMsg + "db find error" + err.Error())
			return err
		}
		if !isExist {
			msg := "记录不存在"
			log.Logger.Error(errorMsg + msg)
			return errors.New(msg)
		}
		tableModelStrBytes, err := json.Marshal(tableModel)
		if err != nil {
			log.Logger.Error(errorMsg + "json Marshal error" + err.Error())
			return err
		}
		rClient.Set(cacheKey, tableModelStrBytes, 12*time.Minute)
	} else if err != nil {
		log.Logger.Error(errorMsg + "redis  Exists  error" + err.Error())
		return err
	} else if isExists > 0 {
		tableModelStr, err := rClient.Get(cacheKey).Result()
		if err != nil {
			log.Logger.Error(errorMsg + "redis  Get  error" + err.Error())
			return err
		}
		err = json.Unmarshal([]byte(tableModelStr), tableModel)
		if err != nil {
			log.Logger.Error(errorMsg + "json Unmarshal error" + err.Error())
			return err
		}
	}
	return nil
}

// FindALlNoCountKvWithCond 针对数据很少的，一般不超过100条数据，整个来存放
func (p *Model) FindAllNoCountKvWithCond(platform string, cacheKey string, rClient *redis.Conn, tableModelList interface{}, whereArgs ...interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	listStr, err := rClient.Get(cacheKey).Result()
	if err == redis.Nil {
		if err := p.FindAllNoCount(platform, tableModelList, whereArgs...); err != nil {
			log.Logger.Error(p.GetTableName() + err.Error())
			return err
		}
		if listBytes, err := json.Marshal(tableModelList); err != nil {
			log.Logger.Error(err.Error())
			return err
		} else { //500*1000*1000
			if err := rClient.Set(cacheKey, string(listBytes), consts.ForeverExpiration).Err(); err != nil {
				log.Logger.Error(err.Error())
				return err
			}
		}
	} else if err != nil {
		return err
	} else {
		err = json.Unmarshal([]byte(listStr), tableModelList)
		if err != nil {
			return err
		}
	}
	return nil
}

// FindALlNoCountKv 针对数据很少的，一般不超过100条数据，整个来存放
func (p *Model) FindAllNoCountKv(platform, cacheKey string, rClient *redis.Conn, tableModelList interface{}) error {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	listStr, err := rClient.Get(cacheKey).Result()
	if err == redis.Nil {
		if err := p.FindAllNoCount(platform, tableModelList); err != nil {
			log.Logger.Error(p.GetTableName() + err.Error())
			return err
		}
		if listBytes, err := json.Marshal(tableModelList); err != nil {
			log.Logger.Error(err.Error())
			return err
		} else {
			//500*1000*1000
			if err := rClient.Set(cacheKey, string(listBytes), 5*time.Minute).Err(); err != nil {
				log.Logger.Error(err.Error())
				return err
			}
		}
	} else if err != nil {
		return err
	} else {
		err = json.Unmarshal([]byte(listStr), tableModelList)
		if err != nil {
			return err
		}

	}
	return nil
}

// GetSession 得到session - 每次拿的都是一个新的session
func (p *Model) GetSession(platform string) *xorm.Session {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	db := p.GetEngine(platform)
	if db == nil {
		panic("无法依据平台识别号获得 db-session")
	}
	return db.NewSession()
}

// GetEngine 得到数据库对象
func (p *Model) GetEngine(platform string) *xorm.EngineGroup {
	if p.GetPlatform != nil {
		platform = p.GetPlatform()
	}
	if d, exists := db.Servers[platform]; exists {
		return d
	}
	return nil
}
