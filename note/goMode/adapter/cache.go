package adapter

type Cache struct {
	localAdapter
}

type localAdapter = Adapter

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) SetAdapter(adapter Adapter) {
	c.localAdapter = adapter
}
