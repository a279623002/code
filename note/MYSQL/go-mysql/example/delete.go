package example

import (
	"fmt"
	"go-mysql/configs"
)

func Delete() {
	r, err := configs.Db.Exec("delete from zb_media where id=?", 160401)
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
