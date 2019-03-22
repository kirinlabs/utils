package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func Type(arg interface{}) string {
	return reflect.TypeOf(arg).String()
}

func Export(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		return ""
	}
	return buf.String()
}

func Json(v interface{}) string {
	return Export(v)
}

func Empty(arg interface{}) bool {
	if arg == nil {
		return true
	}
	v := reflect.ValueOf(arg)
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		if v.Len() > 0 {
			return false
		}
		return true
	case reflect.Ptr:
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func IsInt(arg interface{}) bool {
	switch arg.(type) {
	case int:
		return true
	default:
		return false
	}
}

func IsInt64(arg interface{}) bool {
	switch arg.(type) {
	case int64:
		return true
	default:
		return false
	}
}

func IsFloat64(arg interface{}) bool {
	if Type(arg) == "float64" {
		return true
	}
	return false
}

func IsString(arg interface{}) bool {
	if Type(arg) == "string" {
		return true
	}
	return false
}

func IsUSlice(arg interface{}) bool {
	if Type(arg) == "[]uint8" {
		return true
	}
	return false
}

func IsTime(arg interface{}) bool {
	if Type(arg) == "time.Time" {
		return true
	}
	return false
}

func IsBool(arg interface{}) bool {
	if Type(arg) == "bool" {
		return true
	}
	return false
}

func IsSlice(arg interface{}) bool {
	switch arg.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}
