package main

import (
	"fmt"
	"go-mysql/configs"
)

type NumTest struct {
	Id     int64 `json:"id"`
	Status int64 `json:"status"`
	Type   int64 `json:"type"`
	Num    int64 `json:"num"`
}

func main() {
	mysql := configs.IniMysqlGet()
	configs.InitDb(mysql)

	numTest := []NumTest{}
	err := configs.Db.Select(&numTest, "select * from num_test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numTest)

}
