package cache

import (
	"sync"
	"gin_project/utils/cache/lru"
)

type Cache struct {
    mutex 		sync.Mutex
	lruCache    *lru.LRU
	cacheBytes 	int64
}

// 创建一个构造函数
func NewCache(maxBytes int64) *Cache {
    return &Cache{
    	mutex:      sync.Mutex{},
    	lruCache:   lru.NewLRU(maxBytes, nil),
    	cacheBytes: maxBytes,
    }
}

// 创建一个Get元素方法
func (c *Cache) Get(key string) (value lru.Value, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.lruCache.Get(key)
}

// 创建一个Add元素方法
func (c *Cache) Add(key string, value lru.Value) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.lruCache.Add(key, value)
}