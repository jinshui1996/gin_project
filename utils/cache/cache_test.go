package cache

import (
	"testing"
)

type String string

func (d String) Len() int64 {
	return int64(len(d))
}

var cacheTest = NewCache(100)
func TestCache(t *testing.T) {
    
    // TODO: Implement test cases for Cache
	cacheTest.Add("123", String("123"));
	cacheTest.Add("456", String("456"));
	cacheTest.Add("789", String("789"));
	cacheTest.Get("123");
}