package example

import (
	"fmt"
	"go-mysql/configs"
)

func Update() {
	r, err := configs.Db.Exec("update zb_media set cur_url=? where id=?", "test1", 160401)
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
}
