package consistenthash

import (
	"hash/crc32"
	"log"
	"testing"
)

func TestConsistentHash(t *testing.T) {
    ch := NewConsistentHash(3, crc32.ChecksumIEEE)
	ch.Add("192.168.1.1", "192.168.1.2", "192.168.1.3")
	log.Println(ch.Get("192.168.1.1"))
	log.Println(ch.Get("192.168.1.10"))
}