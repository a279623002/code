package example

import (
	"fmt"
	"go-mysql/configs"
)

func Insert() {
	r, err := configs.Db.Exec("insert into zb_media(cur_url,fid,type,json_data) values(?,?,?,?)", "test", 1, "test", "test")
	if err != nil {
		fmt.Println(err)
		return
	}

	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(row)

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}
