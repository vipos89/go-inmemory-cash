package memory_cash

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	val, exists := c.data[key]
	if !exists {
		return nil
	}
	return val
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
