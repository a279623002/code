package example

import (
	"fmt"
	"go-mysql/configs"
	"math/rand"
	"time"
)

func ToUpdate() {
	conn, err := configs.Db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	name := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < 10000000; i++ {
		sex := 1
		if i % 300 == 0 {
			sex = 2
		}
		num := rand.Intn(len(name))
		r, err := conn.Exec("insert into student (`name`, `sex`) VALUES (?, ?)", name[num], sex)
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
	}

	conn.Commit()
	fmt.Println("done")

}
