package lru

import (
	"container/list"
)

type LRU struct {
	maxBytes int64
	nbytes   int64
	ll		 *list.List
	lru 	 map[string]*list.Element
	// 在从lru lruCache中删除元素是，对该元素的一个回调函数
	OnEvicted func(key string, value Value)
}

// 表示一个缓存项
type entry struct {
    key   string
	value Value
}

// Value接口，用于返回缓存值的大小
type Value interface {
    Len() int64
}

// 实现Value接口的示例
type String string

func (s String) Len() int64 { return int64(len(s)) }

// 定义一个String
var _ Value = String("") // 确保String实现了Value接口

// newLRU 返回一个新的lruCache
func NewLRU(maxBytes int64, onEvicted func(string, Value)) *LRU {
    return &LRU{
		maxBytes: maxBytes,
		ll: list.New(),
		lru: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// add 添加一个值到缓存中
func (c *LRU) Add(key string, value Value) {
    if ele, ok := c.lru[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) // 将缓存项转化为entry类型的指针
		c.nbytes += value.Len() - kv.value.Len()
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.lru[key] = ele
		c.nbytes += int64(value.Len())
	}
	if (c.maxBytes != 0) &&  (c.nbytes > c.maxBytes) {
		c.removeOldest()
	}
}

// removeOldest 移除最旧的缓存项
func (c *LRU) removeOldest() {
    ele := c.ll.Back()
    	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.lru, kv.key)
		c.nbytes -= int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// get 获取缓存中的值
func (c *LRU) Get(key string) (value Value, ok bool) {
    if ele, ok := c.lru[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

// Len 返回缓存中元素的数量
func (c *LRU) Len() int {
    return c.ll.Len()
}