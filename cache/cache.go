package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]*Data
	dmu  sync.RWMutex // data mutex
	list []*Data
	lmu  sync.RWMutex // list mutex
}

type Data struct {
	Key        string
	Value      string
	Expiration int64
	Index      int

	cleanTime *time.Time
}

var localCache *Cache

func InitCache() {
	localCache = &Cache{
		data: make(map[string]*Data),
		dmu:  sync.RWMutex{},
	}
	go func(c *Cache) {
		for {
			select {
			case <-time.Tick(60 * time.Second):
			LABEL:
				for i := range c.list {
					c.lmu.RLock()
					if c.len() < 1 {
						c.lmu.RUnlock()
						break LABEL
					}
					if temp := c.list[i]; temp.Index != 0 && temp.cleanTime.Before(time.Now()) {
						c.lmu.RUnlock()
						c.ttl(temp.Key)
					} else {
						c.lmu.RUnlock()
						break LABEL
					}
				}

			}

		}
	}(localCache)
}

func GetCacheStore() *Cache {
	if localCache == nil {
		InitCache()
	}

	return localCache
}

func (c *Cache) push(data *Data) int {
	c.lmu.Lock()
	defer c.lmu.Unlock()
	c.list = append(c.list, data)
	index := len(c.list) - 1
	return index
}

func (c *Cache) pop() *Data {
	d := &Data{}
	c.lmu.Lock()
	defer c.lmu.Unlock()
	l := len(c.list)
	c.list, d = c.list[:l-1], c.list[l-1]
	return d
}

func (c *Cache) remove(i int) bool {
	c.lmu.Lock()
	defer c.lmu.Unlock()
	c.list = append(c.list[:i], c.list[i+1:]...)
	return true
}

func (c *Cache) len() int {
	return len(c.list)
}

func (c *Cache) ttl(key string) {
	c.dmu.RLock()
	if _, ok := c.data[key]; ok {
		if c.data[key].cleanTime.Before(time.Now()) {
			c.dmu.RUnlock()
			c.delete(key)
		} else {
			c.dmu.RUnlock()
		}
	} else {
		c.dmu.RUnlock()
	}

}

func (c *Cache) get(key string) *Data {
	data := new(Data)
	ok := false
	c.dmu.RLock()
	defer c.dmu.RUnlock()
	if data, ok = c.data[key]; !ok {
		data = nil
	} else if data.Expiration != 0 && data.cleanTime.Before(time.Now()) {
		data = nil
	}
	return data
}

func (c *Cache) set(data *Data) bool {
	//过期时间
	if data.Expiration != 0 {
		ct := time.Now().Add(time.Duration(data.Expiration) * time.Second)
		data.cleanTime = &ct
	}

	c.dmu.Lock()
	defer c.dmu.Unlock()

	data.Index = c.push(data)
	c.data[data.Key] = data

	return true
}

func (c *Cache) delete(key string) bool {
	b := c.get(key)
	if b == nil {
		return false
	}

	c.dmu.Lock()
	defer c.dmu.Unlock()

	c.remove(b.Index)

	delete(c.data, key)

	return true
}

func (c *Cache) GetString(key string) *string {
	result := c.get(key)
	if result == nil {
		return nil
	}

	return &result.Value
}

func (c *Cache) SetString(key string, value string, exp int64) bool {
	data := &Data{
		Key:        key,
		Value:      value,
		Expiration: exp,
		Index:      0,
		cleanTime:  nil,
	}
	return c.set(data)
}

func (c *Cache) RmString(key string) bool {
	return c.delete(key)
}

func (c *Cache) Get(key string) *Data {
	return c.get(key)
}
