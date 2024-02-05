package db

import (
	"sync"
	"time"

	"github.com/go-xorm/xorm"
)

// 保存的数据库session对象集合
type dbSession struct {
	Session  *xorm.Session //session
	Platform string        //平台识别号
	Created  int64         //创建时间
}

// ServerSessions 服务器session
var ServerSessions = struct {
	Sessions map[int]dbSession
	Locker   sync.Mutex
}{
	Sessions: map[int]dbSession{},
	Locker:   sync.Mutex{},
}

// GetSession 获取session
func GetSession(platform string, goID int) *xorm.Session {
	defer ServerSessions.Locker.Unlock()
	ServerSessions.Locker.Lock()

	// 先检测过期的session/超过10秒
	for k, v := range ServerSessions.Sessions {
		if time.Now().Unix()-v.Created > 10 { //如果大于10秒
			if !v.Session.IsClosed() { //如果还没有关闭则关闭之
				v.Session.Close()
			}
			delete(ServerSessions.Sessions, k)
		}
	}

	// 再判断是否存在
	if v, exists := ServerSessions.Sessions[goID]; exists {
		if v.Session.IsClosed() { //如果已经关闭
			v.Session = GetDbByPlatform(platform) //则新建session
			ServerSessions.Sessions[goID] = v     //更新sessions
		}
		return v.Session
	}

	// 如果不存在, 则创建之
	session := GetDbByPlatform(platform)
	defer session.Close()
	ServerSessions.Sessions[goID] = dbSession{
		Session:  session,
		Platform: platform,
		Created:  time.Now().Unix(),
	}

	return session
}
