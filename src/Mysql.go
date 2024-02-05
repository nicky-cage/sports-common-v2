package common

import (
	"sports-common/consts"
	"sports-common/db"

	"github.com/go-xorm/xorm"
)

// Mysql 获得mysql
func Mysql(platform string, masterOnly ...bool) *xorm.Session {
	isMasterOnly := false
	if len(masterOnly) >= 1 {
		isMasterOnly = masterOnly[0]
	}
	return db.GetDbByPlatform(platform, isMasterOnly)
}

// PlatformsMysql 所有平台数据库
func PlatformsMysql() map[string]*xorm.Session {
	arr := map[string]*xorm.Session{}
	for _, platform := range consts.PlatformCodes {
		arr[platform] = Mysql(platform)
	}
	return arr
}

// PlatformsMysqlRestore 关闭所有session
func PlatformsMysqlRestore(allDbSessions map[string]*xorm.Session) {
	for _, dbSession := range allDbSessions {
		dbSession.Close()
	}
}

// UseMysql 使用mysql
func UseMysql(platform string, callback func(*xorm.Session) error) error {
	mConn := Mysql(platform)
	defer mConn.Close()
	return callback(mConn)
}
