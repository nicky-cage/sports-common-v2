package pgsql

import "github.com/go-pg/pg/v10"

// ConnGroup 服务器组
type ConnGroup struct {
	Platform string
	Conns    []*pg.Conn
}

// Close 关闭连接
func (ths *ConnGroup) Close() {
	for _, p := range ths.Conns {
		p.Close()
	}
}
