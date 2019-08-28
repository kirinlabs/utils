// Copyright 2017 The Kirinlabs. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"encoding/json"
	"errors"
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

func Decode(s string) interface{} {
	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		return nil
	}
	return i
}

func Unmarshal(s string, args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("the second parameter is required.")
	}
	err := json.Unmarshal([]byte(s), &args[0])
	if err != nil {
		return err
	}
	return nil
}
