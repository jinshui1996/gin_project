package service

import (
	"gin_project/utils/consistenthash"
)


type HashService struct {
    consistenthash	*consistenthash.ConsistentHash
}

// GetHashKey函数
func (h *HashService) GetHashKey(key string) string {
    return h.consistenthash.Get(key)
}

// AddHashKey函数
func (h *HashService) AddHashKey(key string) {
    h.consistenthash.Add(key)
}

// RemoveHashKey函数
func (h *HashService) RemoveHashKey(key string) {
    h.consistenthash.Remove(key)
}

// ChangeHashKeyNodeNum函数
func (h *HashService) ChangeHashKeyNodeNum(num int) {
    h.consistenthash.SetReplicas(num)
}

// NewHashService 创建hash服务
func NewHashService() *HashService {
    return &HashService{
        consistenthash: consistenthash.NewConsistentHash(10, nil),
    }
}
