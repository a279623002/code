package adapter

type AdapterRedis struct {
	key string
}

func NewAdapterRedis() Adapter {
	return &AdapterRedis{}
}

func (c *AdapterRedis) Set(str string) {
	c.key = str
}

func (c *AdapterRedis) Get() string {
	return c.key
}
