package configs

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sqlx.DB

func InitDb(configs *MysqlConfig) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", configs.UserName, configs.Password, configs.Host, configs.Port, configs.Name))
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db.SetMaxIdleConns(configs.MaxOpenConns)
	db.SetMaxIdleConns(configs.MaxIdleConns)
	Db = db


	//mSql := "select * from user"
	//rows, _ := db.Query(mSql)
	//rows.Close() //这里如果不释放连接到池里，执行5次后其他并发就会阻塞
}
