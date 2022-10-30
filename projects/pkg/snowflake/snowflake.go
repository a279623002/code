package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var node *snowflake.Node

var once sync.Once

//@see:https://github.com/bwmarrin/snowflake
func GetInstance(num int64) {
	once.Do(func() {
		if num == 0 {
			num = 1
		}
		n, err := snowflake.NewNode(num)
		if err != nil {
			fmt.Printf("初始化snowflake失败, err:%s \n", err.Error())
		}
		node = n
	})
}

//发号器
func GetNextIdStr() string {
	return node.Generate().String()
}

func GetNextIdInt() int64 {
	return node.Generate().Int64()
}
