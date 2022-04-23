package example

import (
	"fmt"
	"go-mysql/configs"
)

func ToUpdate() {
	conn, err := configs.Db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := conn.Exec("update zb_media set cur_url=? where id=?", "test1", 160400)
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return
	}
	if row != 1 {
		conn.Rollback()
		fmt.Println("none update")
		return
	}
	conn.Commit()
	fmt.Println("done")

}
