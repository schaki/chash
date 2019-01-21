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

// hasher will provide a hash value for supported types
//nolint: gocyclo
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
	case reflect.Array, reflect.Slice:
		var arrHash uint64
		l := v.Len()
		for i := 0; i < l; i++ {
			t, err := hasher(v.Index(i).Interface())
			if err != nil {
				return 0, err
			}

			if arrHash == 0 {
				arrHash = t
			} else {
				arrHash = arrHash ^ t
			}
		}
		write(h, fmt.Sprintf("%d", arrHash))
	case reflect.Struct:
		var sHash uint64
		f := v.NumField()
		t := v.Type()
		for i := 0; i < f; i++ {
			if t.Field(i).Name == "_" {
				continue
			}
			fieldValue := v.Field(i)
			item, err := hasher(fieldValue.Interface())
			if err != nil {
				return 0, err
			}
			if sHash == 0 {
				sHash = item
			} else {
				sHash = sHash ^ item
			}
		}
		write(h, fmt.Sprintf("%d", sHash))
	default:
		return 0, fmt.Errorf("unsupported kind for hasher %s", k)
	}

	return h.Sum64(), nil
}

// Put stores a value in a Data and gives the client a key
func (d Data) Put(v interface{}) (uint64, error) {
	h, err := hasher(v)
	if err != nil {
		return 0, err
	}
	if d[h] != nil {
		return 0, fmt.Errorf("collision detected hash %d", h)
	}
	d[h] = v
	return h, nil
}

// Get allows for a fast lookup in a Data for content by key
func (d Data) Get(k uint64) interface{} {
	return d[k]
}
