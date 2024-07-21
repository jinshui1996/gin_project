package consistenthash

import (
	"hash/crc32"
	"log"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// 定义一个string的list的实现
type StringList []string

// 定义一个一致性哈希结构
type ConsistentHash struct {
    // TODO: add fields
	hash Hash
	replicas int // 解释
	keys []int
	hashMap map[int]StringList
}

// 添加一个key的实现
func (c *ConsistentHash) Add(keys ...string) {
       for _, key := range keys {
			for i := 0; i < c.replicas; i++ {
				hash := int(c.hash([]byte(strconv.Itoa(i) + key)))
				// 二分查找
				idx := sort.Search(len(c.keys), func(i int) bool {
					return c.keys[i] >= hash
				})
				c.hashMap[hash] = append(c.hashMap[hash], key)
				if idx == len(c.keys) {
					c.keys = append(c.keys, hash)
				} else {
				    return
				}
			}
		}
		sort.Ints(c.keys)
}

// 根据key获取节点
func (c *ConsistentHash) Get(key string) string {
    if len(c.keys) == 0 {
		return ""
	}
	hash := int(c.hash([]byte(key)))
	// 二分查找
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hash
	})
	if idx == len(c.keys) {
		return ""
	}

	// 遍历c.hashMap[c.keys[idx]]的实现
	for i := 0; i < len(c.hashMap[c.keys[idx%len(c.keys)]]); i++ {
		if c.hashMap[c.keys[idx]][i] == key {
			return key
		}
	}
	return ""
}

// 删除一个key的实现
func (c *ConsistentHash) Remove(key string) {
	if len(c.keys) == 0 {
		return
	}
	hash := int(c.hash([]byte(key)))
	// 二分查找
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hash
	})
	if idx == len(c.keys) {
		return
	}

	// 遍历c.hashMap[c.keys[idx]]的实现
	for i := 0; i < len(c.hashMap[c.keys[idx%len(c.keys)]]); i++ {
		if c.hashMap[c.keys[idx]][i] == key {
			c.hashMap[c.keys[idx]] = append(c.hashMap[c.keys[idx]][:i], c.hashMap[c.keys[idx]][i+1:]...)
			return
		}
	}
}

// 修改replicas的实现
func (c *ConsistentHash) SetReplicas(replicas int) {
	c.replicas = replicas
}

// 生成一个ConsistentHash实例
func NewConsistentHash(replicas int, fn Hash) *ConsistentHash {
	if replicas <= 0 {
		log.Fatal("Invalid replicas count")
	}
	if fn == nil {
	    fn = crc32.ChecksumIEEE
	}
    return &ConsistentHash {
			replicas: replicas,
			hash: fn,
			hashMap: make(map[int]StringList),
		}
}