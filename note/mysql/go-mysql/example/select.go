package example

import (
	"fmt"
	"go-mysql/configs"
)

// db table need `db` in tag
type Media struct {
	Id             int64  `json:"id" db:"id"`
	FId            int64  `json:"fid" db:"fid"`
	OriginUrl      string `json:"origin_url" db:"origin_url"`
	OriginBasename string `json:"origin_basename" db:"origin_basename"`
	CurUrl         string `json:"cur_url" db:"cur_url"`
	CurBasename    string `json:"cur_basename" db:"cur_basename"`
	Type           string `json:"type" db:"type"`
	JsonData       string `json:"json_data" db:"json_data"`
}

func Select() {
	res := []Media{}
	err := configs.Db.Select(&res, "select * from zb_media")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

