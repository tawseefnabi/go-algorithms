package hashmap

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   any
	value any
	next  *node
}

// HashMap is golang implementation of hashmap
type HashMap struct {
	size     uint64
	capacity uint64
	table    []*node
}

func New() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity, defaultCapacity),
	}
}

// Make creates a new HashMap instance with input size and capacity
func Make(size, capacity uint64) HashMap {
	return HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, defaultCapacity, capacity),
	}
}
func (hm *HashMap) Get() {
	fmt.Println("--------")
}

// Put puts new key value in hashmap
func (hm *HashMap) Put(key any, value any) any {
	fmt.Println("put", key, value, defaultCapacity)
	fmt.Println("0000", hm.hash(key))
	hm.putValue(hm.hash(key), key, value)
	return "kk"
}

func (hm *HashMap) putValue(hash uint64, key any, value any) any {
	if hm.capacity == 0 {
		hm.capacity = defaultCapacity
		hm.table = make([]*node, defaultCapacity)
	}
	fmt.Println("hm.table[hash]", hash, key)
	// fmt.Println(hm.table[11])
	node := hm.getNodeByHash(hash)
	if node == nil {
		fmt.Println("0000")
		// hm.table[hash] = newNode(key, value)

	} else if node.key == key {
		fmt.Println("====")
		// hm.table[hash] = newNodeWithNext(key, value, node)
		// return value

	} else {
		fmt.Println("-----")
		// hm.resize()
		// return hm.putValue(hash, key, value)
	}
	return "lll"
}

func (hm *HashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := h.Sum64()
	return hashValue
}
