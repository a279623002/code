package example

import (
	"fmt"
	"my-redis/configs"
	"my-redis/utils"
)

func PrintHash() {
	fmt.Println(utils.JSHash("shiro"))
	fmt.Println(utils.RSHash("shiro"))
	fmt.Println(utils.PJWHash("shiro"))
	fmt.Println(utils.BKDRHash("shiro"))
	fmt.Println(utils.SDBMHash("shiro"))
	fmt.Println(utils.DJBHash("shiro"))
	fmt.Println(utils.DEKHash("shiro"))
	fmt.Println(utils.APHash("shiro"))
}

type BloomFilter struct {
	Bucket    string
	HashFuncs []string
}

//Redis BitMap 的底层数据结构实际上是 String 类型，Redis 对于 String 类型有最大值限制不得超过 512M，即 2^32 次方 byte
//int64 2^63-1 ERR bit offset is not an integer or out of range
func NewBloomFilter(bucket string, hashFuncs []string) *BloomFilter {
	return &BloomFilter{
		Bucket:    bucket,
		HashFuncs: hashFuncs,
	}
}

func (b *BloomFilter) Add(key string) {
	for _, v := range b.HashFuncs {
		hash := int64(utils.GetHash(v, key))

		_, err := configs.RedisDB.SetBit(b.Bucket, hash, 1).Result()
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

func (b *BloomFilter) Exists(key string) bool {
	res := []int64{}
	for _, v := range b.HashFuncs {
		hash := int64(utils.GetHash(v, key))
		if hash < 0 {
			hash *= -1
		}
		item, err := configs.RedisDB.GetBit(b.Bucket, hash).Result()
		if err != nil {
			return false
		}
		res = append(res, item)
	}
	for _, v := range res {
		if v == 0 {
			return false
		}
	}
	return true

}
