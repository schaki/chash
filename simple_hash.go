package simplehash

import (
	"fmt"
	"hash/fnv"
)

// Data is a structure to store hash-keyed content
type Data map[uint64]string

func hash(s string) (uint64, error) {
	h := fnv.New64()
	h.Reset()
	_, err := h.Write([]byte(s))
	if err != nil {
		fmt.Printf("there was an error writing the string\n")
		return 0, err
	}
	return h.Sum64(), nil
}

// Put stores a value in a Data and gives the user a key
func (d Data) Put(v string) (uint64, error) {
	h, err := hash(v)
	if err != nil {
		return 0, err
	}
	d[h] = v
	return h, nil
}

// Get allows for a fast lookup in a Data for content by key
func (d Data) Get(k uint64) string {
	return d[k]
}
