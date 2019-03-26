package utils

import (
	"testing"

	"utils/sys"
)

func TestRealPath(t *testing.T) {
	filename := "./utils.go"
	sys.RealPath(filename)
}

func TestEmpty(t *testing.T) {
	type KV struct {
		value  interface{}
		result bool
	}

	data := map[string]*KV{
		"int_true":     &KV{0, true},
		"int_false":    &KV{100, false},
		"str_true":     &KV{"", true},
		"str_false":    &KV{"github", false},
		"slice_true":   &KV{make([]int, 0), true},
		"slice_false":  &KV{[]int{1, 2}, false},
		"map_true":     &KV{map[string]interface{}{}, true},
		"map_false":    &KV{map[string]interface{}{"name": "utils"}, false},
		"ptr_true":     &KV{&KV{}, true},
		"ptr_false":    &KV{&KV{"utils", false}, false},
		"struct_true":  &KV{KV{}, true},
		"struct_false": &KV{KV{"utils", false}, false},
	}

	for k, v := range data {
		if Empty(v.value) != v.result {
			t.Errorf("Empty test failed: %s expected %v,return %v != %v", k, v.result, Empty(v.value), v.result)
		}
	}
}
