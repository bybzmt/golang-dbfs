package dbfs

import (
	"sync"
	"github.com/golang/groupcache/lru"
)

type Cache struct {
	L sync.Mutex
	lru.Cache
}

func (c *Cache) Init(MaxEntries int) {
	c.Cache = lru.New(MaxEntries)
}

func (c *Cache) Add(key lru.Key, value interface{}) {
	c.L.Lock()
	defer c.L.Unlock()

	c.Cache.Add(key, value)
}

func (c *Cache) Get(key lru.Key) (value interface{}, ok bool) {
	c.L.Lock()
	defer c.L.Unlock()

	return c.Cache.Get(key)
}

func (c *Cache) Remove(key lru.Key) {
	c.L.Lock()
	defer c.L.Unlock()

	c.Remove(key)
}

//----具体的cache-----------

type nodeCache struct {
	Cache
}

func (c *nodeCache) Add(key int64, val *RowNode) {
	c.Cache.Add(key, val)
}

func (c *nodeCache) Get(key int64) (value *RowNode, ok bool) {
	val, ok := c.Cache.Get(key)
	if ok {
		return val.(*RowNode), ok
	}
	return nil, ok
}

func (c *nodeCache) Remove(key int64) {
	c.Cache.Remove(key)
}

//-----------

type fileCache struct {
	Cache
}

func (c *fileCache) Add(key int64, val *RowFile) {
	c.Cache.Add(key, val)
}

func (c *fileCache) Get(key int64) (value *RowFile, ok bool) {
	val, ok := c.Cache.Get(key)
	if ok {
		return val.(*RowFile), ok
	}
	return nil, ok
}

func (c *fileCache) Remove(key int64) {
	c.Cache.Remove(key)
}

//-----------

type storageCache struct {
	Cache
}

func (c *storageCache) Add(key int64, val *RowStorage) {
	c.Cache.Add(key, val)
}

func (c *storageCache) Get(key int64) (value *RowStorage, ok bool) {
	val, ok := c.Cache.Get(key)
	if ok {
		return val.(*RowStorage), ok
	}
	return nil, ok
}

func (c *storageCache) Remove(key int64) {
	c.Cache.Remove(key)
}

