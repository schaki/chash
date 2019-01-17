package simplehash

import (
	"fmt"
	"hash"
	"hash/fnv"
	"reflect"
)

// Data is a structure to store hash-keyed content
type Data map[uint64]interface{}

func write(h hash.Hash64, s string) {
	h.Reset()

	_, err := h.Write([]byte(s))
	if err != nil {
		fmt.Printf("there was an error writing the string\n")
	}
}

func hasher(i interface{}) (uint64, error) {
	h := fnv.New64()

	v := reflect.ValueOf(i)

	k := v.Kind()
	switch k {
	case reflect.Int:
		v = reflect.ValueOf(v.Int())
		write(h, fmt.Sprintf("%d", v))
	case reflect.Uint:
		v = reflect.ValueOf(v.Uint())
		write(h, fmt.Sprintf("%d", v))
	case reflect.String:
		write(h, v.String())
	}

	if k == reflect.Int || k == reflect.Uint || k == reflect.String {
		return h.Sum64(), nil
	}
	return 0, fmt.Errorf("unsupported kind for hasher")
}

// Put stores a value in a Data and gives the user a key
func (d Data) Put(v interface{}) (uint64, error) {
	h, err := hasher(v)
	if err != nil {
		return 0, err
	}
	d[h] = v
	return h, nil
}

// Get allows for a fast lookup in a Data for content by key
func (d Data) Get(k uint64) interface{} {
	return d[k]
}
