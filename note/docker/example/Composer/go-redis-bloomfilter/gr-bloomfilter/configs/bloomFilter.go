package configs

import (
	"context"
	"gr-bloomfilter/utils"
)

type BloomFilter struct {
	Bucket    string
	HashFuncs []string
}

func (b *BloomFilter) Add(ctx context.Context, key string) (err error) {
	for _, v := range b.HashFuncs {
		hash := int64(utils.GetHash(v, key))

		_, err := RedisDB.SetBit(b.Bucket, hash, 1).Result()
		if err != nil {
			return err
		}

	}
	return nil
}

func (b *BloomFilter) Exists(ctx context.Context, key string) (res bool, err error) {
	target := []int64{}
	for _, v := range b.HashFuncs {
		hash := int64(utils.GetHash(v, key))
		if hash < 0 {
			hash *= -1
		}
		item, err := RedisDB.GetBit(b.Bucket, hash).Result()
		if err != nil {
			return false, err
		}
		target = append(target, item)
	}
	for _, v := range target {
		if v == 0 {
			return false, nil
		}
	}
	return true, nil

}
