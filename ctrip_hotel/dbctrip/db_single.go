package dbctrip

var server *Mysqlserver

func GetInstance() *Mysqlserver {
	if server == nil {
		server, err := NewMysqlserver()
		if err != nil {
			panic(nil)
		}
		return server
	}
	return server
}
