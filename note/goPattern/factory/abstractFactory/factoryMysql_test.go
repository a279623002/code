package abstractFactory

import "testing"

func TestMysqlFactory_CreateDBConnection(t *testing.T) {
	var db DBFactory
	db = &MysqlFactory{}
	if res := db.CreateDBConnection().Connection(); res != "mysql connect" {
		t.Errorf("res expected be mysql connect, but %s got", res)
	}
}
