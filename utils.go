// Copyright 2017 The Kirinlabs. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utils implements a simple public method for rapid development.

package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
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

// Returns a Json string for the variable
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

// Returns a Json string for the variable
func Json(v interface{}) string {
	return Export(v)
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
