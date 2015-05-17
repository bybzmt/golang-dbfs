package dbfs

import (
	"github.com/golang/groupcache/lru"
)

type nodeCache struct {
	lru.Cache
}

func (c *nodeCache) Add(key int64, value *RowNode) {
	c.Cache.Add(key, value)
}

func (c *nodeCache) Get(key int64) (value *RowNode, ok bool) {
	val, ok := c.Get(key)
	if ok {
		return val.(*RowNode), ok
	}

	return nil, ok
}

func (c *nodeCache) Remove(key int64) {
	c.Remove(key)
}

type storageCache struct {
	lru.Cache
}

func (c *storageCache) Add(key int64, value *RowStorage) {
	c.Cache.Add(key, value)
}

func (c *storageCache) Get(key int64) (value *RowStorage, ok bool) {
	val, ok := c.Get(key)
	if ok {
		return val.(*RowNode), ok
	}

	return nil, ok
}

func (c *storageCache) Remove(key int64) {
	c.Remove(key)
}
