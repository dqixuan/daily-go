package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

func main() {
	
}

type ConsistentHash struct {
	virtualNodes []uint32
	virtualNodeMap map[uint32]string
	nodes map[string]bool
	virtualCount int32
	lock sync.Mutex
}

func (c *ConsistentHash) hashKey(node string) uint32 {
	return crc32.ChecksumIEEE([]byte(node))
}

func (c *ConsistentHash) AddNode(node string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if node == "" {
		return errors.New("node can not be empty.")
	}
	if c.virtualNodeMap == nil {
		c.virtualNodeMap = make(map[uint32]string)
	}
	if c.nodes == nil {
		c.nodes = make(map[string]bool)
	}
	if _, ok := c.nodes[node]; ok {
		return errors.New("node already exists")
	}
	c.nodes[node] = true
	for i:=0; i < int(c.virtualCount); i++ {
		hashKey := c.hashKey(fmt.Sprintf("%s%d", node, i))
		c.virtualNodes = append(c.virtualNodes, hashKey)
		c.virtualNodeMap[hashKey] = node
	}
	return nil
}

func (c *ConsistentHash) search(hashKey uint32) int {
	idx := sort.Search(len(c.virtualNodes), func(i int) bool {
		return c.virtualNodes[i] >= hashKey
	})
	if idx < len(c.virtualNodes) {
		if idx == len(c.virtualNodes)-1 {
			return 0
		}
		return idx
	} else {
		return len(c.virtualNodes)-1
	}
}

func (c *ConsistentHash) Get(node string) string {
	idx := c.search(c.hashKey(node))
	return c.virtualNodeMap[c.virtualNodes[idx]]
}

func (c *ConsistentHash) Delete(node string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	return nil
}
