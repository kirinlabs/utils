// Copyright 2017 The Kirinlabs. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utils implements a simple public method for rapid development.

package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

// Determine if the variable is null
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

// Returns a formatted Json string
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

// Returns a Json string
func Json(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// Returns a interface{}
func Decode(s string, args ...interface{}) interface{} {
	var i interface{}
	if len(args) > 0 {
		i = args[0]
	}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		return nil
	}
	return i
}

// Returns a string of variable type
func Type(arg interface{}) string {
	return reflect.TypeOf(arg).String()
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
	if reflect.TypeOf(arg).Kind().String() == "slice" {
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

func JoinInt(s []int, sep ...string) string {
	sepLetter := ","
	if len(sep) > 0 {
		sepLetter = sep[0]
	}
	str := make([]string, 0)
	for _, v := range s {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, sepLetter)
}

func SplitInt(str string, sep ...string) ([]int, error) {
	sepLetter := ","
	if len(sep) > 0 {
		sepLetter = sep[0]
	}
	i := make([]int, 0)
	s := strings.Split(str, sepLetter)
	for _, v := range s {
		d, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		i = append(i, d)
	}
	return i, nil
}
