package sli

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/kirinlabs/utils/str"
)

func InSlice(v interface{}, s interface{}) bool {
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() == reflect.Slice {
		log.Println("Warning: First parameter cannot be a slice, FUNCTION: InSlice")
		return false
	}

	sv := reflect.Indirect(reflect.ValueOf(s))
	if sv.Kind() != reflect.Slice {
		log.Println("Warning: Second parameter Needs a slice, FUNCTION: InSlice")
		return false
	}

	st := reflect.TypeOf(s).Elem().String()
	vt := reflect.TypeOf(v).String()
	if st != vt {
		log.Println("Warning: Two parameters are not the same type, FUNCTION: InSlice")
		return false
	}

	switch vt {
	case "string":
		for _, vv := range s.([]string) {
			if vv == v {
				return true
			}
		}
	case "int":
		for _, vv := range s.([]int) {
			if vv == v {
				return true
			}
		}
	case "int64":
		for _, vv := range s.([]int64) {
			if vv == v {
				return true
			}
		}
	case "float64":
		for _, vv := range s.([]float64) {
			if vv == v {
				return true
			}
		}
	default:
		log.Println("Warning: This type is not supported, FUNCTION: InSlice")
		return false
	}

	return false
}

func InIfaceSlice(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func Slice(iface interface{}) ([]interface{}, error) {
	d := make([]interface{}, 0)
	v := reflect.Indirect(reflect.ValueOf(iface))
	if v.Kind() != reflect.Slice {
		return []interface{}{}, errors.New("Needs a slice")
	}
	t := reflect.TypeOf(iface).Elem().String()
	switch t {
	case "string":
		for _, v := range iface.([]string) {
			if strings.Trim(str.String(v), " ") == "" {
				continue
			}
			d = append(d, v)
		}
	case "int":
		for _, v := range iface.([]int) {
			d = append(d, v)
		}
	case "int64":
		for _, v := range iface.([]int64) {
			d = append(d, v)
		}
	case "float64":
		for _, v := range iface.([]float64) {
			d = append(d, v)
		}
	default:
		return d, errors.New("This type is not supported")
	}
	return d, nil
}

func Unique(list *[]string) []string {
	r := make([]string, 0)
	temp := map[string]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

func UniqueInt(list *[]int) []int {
	r := make([]int, 0)
	temp := map[int]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

/*func UniqueInt(list *[]int) []int {
	var r []int = []int{}
	for _, i := range *list {
		if len(r) == 0 {
			r = append(r, i)
		} else {
			for k, v := range r {
				if i == v {
					break
				}
				if k == len(r)-1 {
					r = append(r, i)
				}
			}
		}
	}
	return r
}*/

func UniqueInt64(list *[]int64) []int64 {
	r := make([]int64, 0)
	temp := map[int64]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

func Chunk(slice []interface{}, size int) (chunk [][]interface{}) {
	if size <= 0 {
		chunk = make([][]interface{}, 0)
		return
	}
	if size >= len(slice) {
		chunk = append(chunk, slice)
		return
	}
	end := size
	for i := 0; i < len(slice); i += size {
		if end > len(slice) {
			chunk = append(chunk, slice[i:])
			break
		}
		chunk = append(chunk, slice[i:end])
		end += size
	}
	return
}

func Rand(a []string) (b string) {
	r := rand.Intn(len(a))
	b = a[r]
	return
}

func RandInt(a []int) (b int) {
	r := rand.Intn(len(a))
	b = a[r]
	return
}

func RandInt64(a []int64) (b int64) {
	r := rand.Intn(len(a))
	b = a[r]
	return
}

func RandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

func Sum(i []int64) (s int64) {
	for _, v := range i {
		s += v
	}
	return
}

func Diff(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func DiffInt(slice1, slice2 []int) (diffslice []int) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func DiffInt64(slice1, slice2 []int64) (diffslice []int64) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func Intersect(slice1, slice2 []string) (interSlice []string) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			interSlice = append(interSlice, v)
		}
	}
	return
}

func IntersectInt(slice1, slice2 []int) (interSlice []int) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			interSlice = append(interSlice, v)
		}
	}
	return
}

func IntersectIn64(slice1, slice2 []int64) (interSlice []int64) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			interSlice = append(interSlice, v)
		}
	}
	return
}

func Range(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

func Pad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InIfaceSlice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

func Shuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
