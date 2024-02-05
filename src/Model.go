package common

import (
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"xorm.io/builder"
)

// IModel 模型接口
type IModel interface {
	CreateAll(string, []map[string]interface{}) error                                        // 参数: 平台识别号/提交数据
	FindById(string, int, interface{}) (bool, error)                                         // 依据平台/编号来获取记录
	FindAll(string, interface{}, ...interface{}) (uint64, error)                             // 参数: 平台/出参/builder.Cond/limit/offset/orderBy
	FindAllNoCount(string, interface{}, ...interface{}) error                                // 参数: 平台/出参/builder.Cond/limit/offset/orderBy
	Find(string, interface{}, ...interface{}) (bool, error)                                  // 参数: 平台/出参/builder.Cond/orderBy
	Create(string, map[string]interface{}) (uint64, error)                                   // 参数: 平台/数据
	Update(string, map[string]interface{}) error                                             // 参数: 平台/编号/数据
	Delete(string, string) error                                                             // 参数: 平台/编号
	GetEngine(string) *xorm.EngineGroup                                                      // 获取数据库连接
	FindByIdHash(string, string, int, *redis.Conn, interface{}) error                        // 通过id获取hash缓存并转换成struct 只支持一级
	FindByIdKv(string, string, int, bool, *redis.Conn, interface{}) error                    // 通过id获取key-value缓存并转换成struct 只支持一级
	FindAllNoCountKv(string, string, *redis.Conn, interface{}) error                         // 通过id获取key-value缓存并转换成struct 只支持一级
	FindByAndWhereKv(string, string, *redis.Conn, interface{}, string, ...interface{}) error // 通过id获取key-value缓存并转换成struct 只支持一级
	GetFindParameters(args ...interface{}) (builder.Cond, int, int, string)                  // 获取分页、查询条件、offset、order-by
	GetSession(string) *xorm.Session                                                         // 获取session - 每次都是最新的
	GetTableName() string                                                                    // 获取表名称
}

// Model 模型基类
type Model struct {
	TabName     string        `json:"-"` // 表名称
	GetPlatform func() string `json:"-"` // 获取平台名称
	IsFrontend  bool          `json:"-"` // 是否只用于前台
	CacheLevel  int           `json:"-"` // 0 不缓存， 1缓存
	CachePrefix string        `json:"-"` // 缓存的key构造
	Joins       [][]string    `json:"-"` // joins [][]string{{"left", "users", "users.id = user_accounts.id"}}

	*xorm.Session `json:"-"` // session - 仅用于继承原有方法
}
